package service

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/SaenkoDmitry/training-tg-bot/internal/constants"
	"github.com/SaenkoDmitry/training-tg-bot/internal/messages"
	summarysvc "github.com/SaenkoDmitry/training-tg-bot/internal/service/summary"
	"github.com/SaenkoDmitry/training-tg-bot/internal/service/tghelpers"
	"strconv"
	"strings"
	"time"

	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/users"

	"github.com/SaenkoDmitry/training-tg-bot/internal/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *serviceImpl) HandleMessage(message *tgbotapi.Message) {
	chatID := message.Chat.ID
	text := message.Text

	user, _ := s.usersRepo.GetByChatID(chatID)

	fmt.Println("HandleMessage:", text)

	switch {
	case text == messages.BackToMenu || text == "/start" || text == "/menu":
		s.sendMainMenu(chatID, message.From)

	case text == messages.StartWorkout || text == "/start_workout":
		s.showWorkoutTypeMenu(chatID)

	case text == messages.MyWorkouts || text == "/workouts":
		s.showMyWorkouts(chatID, 0)

	case text == messages.Stats || text == "/stats":
		s.showStatsMenu(chatID)

	case text == messages.Settings || text == "/settings":
		s.settings(chatID)

	case text == messages.HowToUse || text == "/about":
		s.about(chatID)

	case text == messages.Admin || text == "/admin":
		s.admin(chatID, user)

	default:
		s.handleState(chatID, text)
	}
}

func (s *serviceImpl) sendMainMenu(chatID int64, from *tgbotapi.User) {
	method := "sendMainMenu"

	text := messages.Hello

	user := s.createUserIfNotExists(chatID, from)

	rows := make([][]tgbotapi.KeyboardButton, 0)
	rows = append(rows, tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(messages.StartWorkout),
	))
	rows = append(rows, tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(messages.MyWorkouts),
		tgbotapi.NewKeyboardButton(messages.Stats),
	))
	rows = append(rows, tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(messages.Settings),
		tgbotapi.NewKeyboardButton(messages.HowToUse),
	))

	if user.IsAdmin() {
		rows = append(rows, tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(messages.Admin),
		))
	}

	keyboard := tgbotapi.NewReplyKeyboard(rows...)
	keyboard.ResizeKeyboard = true

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = constants.HtmlParseMode
	msg.ReplyMarkup = keyboard
	_, err := s.bot.Send(msg)
	handleErr(method, err)
}

func (s *serviceImpl) createUserIfNotExists(chatID int64, from *tgbotapi.User) *models.User {
	user, err := s.usersRepo.GetByChatID(chatID)
	if err == nil {
		return user
	}

	if errors.Is(err, users.NotFoundUserErr) {
		createdUser, createErr := s.usersRepo.Create(chatID, from)
		if createErr != nil {
			return nil
		}

		// создаем дефолтную программу
		program, createErr := s.programsRepo.Create(createdUser.ID, "#1 стартовая")
		if createErr != nil {
			return nil
		}

		// прикрепляем программу к юзеру и сохраняем
		createdUser.ActiveProgramID = &program.ID
		err = s.usersRepo.Save(createdUser)
		if err != nil {
			return nil
		}
		return createdUser
	}
	return nil
}

func (s *serviceImpl) showWorkoutTypeMenu(chatID int64) {
	method := "showWorkoutTypeMenu"

	user, err := s.usersRepo.GetByChatID(chatID)
	if err != nil {
		s.handleGetUserErr(chatID, method, err)
		return
	}

	program, err := s.programsRepo.Get(*user.ActiveProgramID)
	if err != nil {
		return
	}

	if len(program.DayTypes) == 0 {
		msg := tgbotapi.NewMessage(chatID, "Добавьте тренировочные дни в программу через '⚙️ Настройки'")
		msg.ParseMode = constants.MarkdownParseMode
		_, err = s.bot.Send(msg)
		handleErr(method, err)
		return
	}

	text := "*Выберите день тренировки:*"

	buttons := make([][]tgbotapi.InlineKeyboardButton, 0)

	for i, day := range program.DayTypes {
		if i%2 == 0 {
			buttons = append(buttons, []tgbotapi.InlineKeyboardButton{})
		}
		buttons[len(buttons)-1] = append(buttons[len(buttons)-1],
			tgbotapi.NewInlineKeyboardButtonData(day.Name, fmt.Sprintf("workout_create_%d", day.ID)),
		)
	}
	buttons = append(buttons, []tgbotapi.InlineKeyboardButton{})

	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons...)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard
	msg.ParseMode = constants.MarkdownParseMode
	_, err = s.bot.Send(msg)
	handleErr(method, err)
}

func (s *serviceImpl) handleGetUserErr(chatID int64, method string, err error) {
	if errors.Is(err, users.NotFoundUserErr) {
		msg := tgbotapi.NewMessage(chatID, "Сначала создайте пользователя в боте, через команду /start")
		_, err = s.bot.Send(msg)
		handleErr(method, err)
	}
}

func handleErr(method string, err error) {
	if err != nil {
		fmt.Printf("\n %s: error is: %s \n", method, err.Error())
	}
}

func (s *serviceImpl) showWorkoutsByUser(chatID, userID int64) {
	method := "showWorkoutsByUser"
	workouts, err := s.workoutsRepo.FindAll(userID)
	if err != nil {
		return
	}

	user, err := s.usersRepo.GetByID(userID)
	if err != nil {
		return
	}

	if len(workouts) == 0 {
		msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("📭 У пользователя %s пока нет созданных тренировок.", user.ShortName()))
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(messages.BackToMenu, "back_to_menu"),
			),
		)
		msg.ReplyMarkup = keyboard
		_, _ = tghelpers.SendMessage(s.bot, msg, method)
		return
	}

	text := fmt.Sprintf("📋 <b>Тренировки пользователя '%s'</b>\n\n", user.ShortName())
	for i, workout := range workouts {
		status := "🟡"
		if workout.Completed {
			status = "✅"
			if workout.EndedAt != nil {
				status += fmt.Sprintf(" ~ %s",
					utils.BetweenTimes(workout.StartedAt, workout.EndedAt),
				)
			}
		}
		date := workout.StartedAt.Add(3 * time.Hour).Format("02.01.2006 15:04")

		dayType := workout.WorkoutDayType

		text += fmt.Sprintf("%d. <b>%s</b> %s\n   📅 %s\n\n",
			i+1, dayType.Name, status, date)
	}

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = constants.HtmlParseMode
	_, _ = tghelpers.SendMessage(s.bot, msg, method)
}

const (
	showWorkoutsLimit = 4
)

func (s *serviceImpl) showMyWorkouts(chatID int64, offset int) {
	method := "showMyWorkouts"
	user, err := s.usersRepo.GetByChatID(chatID)
	if err != nil {
		s.handleGetUserErr(chatID, method, err)
		return
	}

	count, _ := s.workoutsRepo.Count(user.ID)

	limit := showWorkoutsLimit

	workouts, _ := s.workoutsRepo.Find(user.ID, offset, limit)

	if len(workouts) == 0 {
		msg := tgbotapi.NewMessage(chatID, "📭 У вас пока нет созданных тренировок.\n\nСоздайте первую тренировку!")
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(messages.BackToMenu, "back_to_menu"),
			),
		)
		msg.ReplyMarkup = keyboard
		_, _ = tghelpers.SendMessage(s.bot, msg, method)
		return
	}

	var rows [][]tgbotapi.InlineKeyboardButton

	text := fmt.Sprintf("📋 <b>Ваши тренировки (%d-%d из %d):</b>\n\n", offset+1, min(offset+limit, int(count)), count)
	for i, workout := range workouts {
		status := "🟡"
		if workout.Completed {
			status = "✅"
			if workout.EndedAt != nil {
				status += fmt.Sprintf(" ~ %s",
					utils.BetweenTimes(workout.StartedAt, workout.EndedAt),
				)
			}
		}
		date := workout.StartedAt.Add(3 * time.Hour).Format("02.01.2006 в 15:04")

		dayType := workout.WorkoutDayType

		text += fmt.Sprintf("%d. <b>%s</b> %s\n   📅 %s\n\n",
			i+1+offset, dayType.Name, status, date)

		// buttons
		if i%2 == 0 {
			rows = append(rows, []tgbotapi.InlineKeyboardButton{})
		}
		rows[len(rows)-1] = append(rows[len(rows)-1],
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s %d", dayType.Name, i+1+offset),
				fmt.Sprintf("workout_show_progress_%d", workout.ID)))
	}

	text += "<b>Выберите тренировку для просмотра:</b>"

	rows = append(rows, []tgbotapi.InlineKeyboardButton{})
	fmt.Println("offset", offset, "limit", limit, "count", count)
	if offset >= limit {
		rows[len(rows)-1] = append(rows[len(rows)-1], tgbotapi.NewInlineKeyboardButtonData("⬅️ Предыдущие",
			fmt.Sprintf("workout_show_my_%d", offset-limit)))
	}
	if offset+limit < int(count) {
		rows[len(rows)-1] = append(rows[len(rows)-1], tgbotapi.NewInlineKeyboardButtonData("➡️ Следующие",
			fmt.Sprintf("workout_show_my_%d", offset+limit)))
	} else {
		rows = append(rows, []tgbotapi.InlineKeyboardButton{})
		rows[len(rows)-1] = append(rows[len(rows)-1], tgbotapi.NewInlineKeyboardButtonData("🔙 В начало", "workout_show_my"))
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = constants.HtmlParseMode
	msg.ReplyMarkup = keyboard
	_, _ = tghelpers.SendMessage(s.bot, msg, method)
}

func (s *serviceImpl) showStatsMenu(chatID int64) {
	method := "showStatsMenu"
	text := "📊 *Статистика тренировок*\n\n Выберите период:"

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📅 За неделю", "stats_week"),
			tgbotapi.NewInlineKeyboardButtonData("🗓️ За месяц", "stats_month"),
			tgbotapi.NewInlineKeyboardButtonData("📈 Общая", "stats_all"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = constants.MarkdownParseMode
	msg.ReplyMarkup = keyboard
	_, err := s.bot.Send(msg)
	handleErr(method, err)
}

func (s *serviceImpl) settings(chatID int64) {
	method := "settings"
	buttons := make([][]tgbotapi.InlineKeyboardButton, 0)
	buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(messages.ProgramManagement, "program_management"),
	))
	buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(messages.Export, "export_to_excel"),
	))
	msg := tghelpers.NewMessageBuilder().WithChatID(chatID).WithText("<b>Выберите действие:</b>").WithReplyMarkup(buttons).Build()
	_, _ = tghelpers.SendMessage(s.bot, msg, method)
}

func (s *serviceImpl) users(chatID int64, user *models.User) {
	if !user.IsAdmin() {
		return
	}
	method := "users"
	rows := make([][]tgbotapi.InlineKeyboardButton, 0)

	userObjs, err := s.usersRepo.GetTop10()
	if err != nil {
		return
	}
	var text bytes.Buffer
	text.WriteString(fmt.Sprintf("<b>%s:</b>\n\n", messages.Users))
	for i, u := range userObjs {
		if i%2 == 0 {
			rows = append(rows, tgbotapi.NewInlineKeyboardRow())
		}
		text.WriteString(fmt.Sprintf("• %s\n\n", u.FullName()))
		rows[len(rows)-1] = append(rows[len(rows)-1],
			tgbotapi.NewInlineKeyboardButtonData(u.Username, fmt.Sprintf("workout_show_by_user_id_%d", u.ID)),
		)
	}
	msg := tghelpers.NewMessageBuilder().WithChatID(chatID).WithText(text.String()).WithReplyMarkup(rows).Build()
	_, _ = tghelpers.SendMessage(s.bot, msg, method)
}

func (s *serviceImpl) admin(chatID int64, user *models.User) {
	if !user.IsAdmin() {
		return
	}
	method := "admin"
	rows := make([][]tgbotapi.InlineKeyboardButton, 0)
	rows = append(rows, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(messages.Users, "/admin/users"),
	))
	msg := tghelpers.NewMessageBuilder().WithChatID(chatID).WithText("<b>👨🏻‍💻 Админ панель</b>").WithReplyMarkup(rows).Build()
	_, _ = tghelpers.SendMessage(s.bot, msg, method)
}

func (s *serviceImpl) about(chatID int64) {
	method := "about"
	msg := tgbotapi.NewMessage(chatID, `
	<b>Form Journey · Training 🏔️</b>

	Трекер тренировок и прогресса 📈  
	Измеримый путь к лучшей форме через систему и процесс.
	
	<b>Возможности бота:</b>
	
	• 📕 Персональные программы тренировок  
	• ✍️ Учёт весов, повторов, времени и дистанций  
	• 🤓 Сохранение прошлых показателей для следующей тренировки  
	• 🎥 Видео с техникой выполнения упражнений  
	• ⏱️ Таймеры отдыха между подходами  
	• 📊 Статистика и аналитика тренировок  
	
	<b>Основные разделы:</b>
	
	▶️ <b>Начать тренировку</b> — выполнение текущей программы  
	📋 <b>Мои тренировки</b> — история тренировок  
	📊 <b>Статистика</b> — сводка и динамика прогресса  
	⚙️ <b>Настройки</b> — управление программами и упражнениями + экспорт данных в Excel
	`)

	msg.ParseMode = constants.HtmlParseMode
	_, err := s.bot.Send(msg)
	handleErr(method, err)
}

func (s *serviceImpl) handleState(chatID int64, text string) {
	method := "handleState"
	state, exists := s.userStatesMachine.GetValue(chatID)
	if !exists {
		return
	}

	var err error

	switch {
	case strings.HasPrefix(state, "awaiting_reps_"):
		err = s.awaitingEnterData(
			chatID, state,
			func() (interface{}, error) { return strconv.ParseInt(text, 10, 64) },
			func(nextSet models.Set, value interface{}) models.Set {
				reps, ok := value.(int64)
				if !ok {
					return models.Set{}
				}
				nextSet.FactReps = int(reps)
				if int(reps) != nextSet.Reps {
					nextSet.FactReps = int(reps)
				} else {
					nextSet.FactReps = 0
				}
				return nextSet
			},
			"❌ Неверный формат числа повторений. Введите целое число (например: 42)",
			"✅ Количество повторений обновлено",
		)
	case strings.HasPrefix(state, "awaiting_weight_"):
		err = s.awaitingEnterData(
			chatID, state,
			func() (interface{}, error) { return strconv.ParseFloat(text, 32) },
			func(nextSet models.Set, value interface{}) models.Set {
				weight, ok := value.(float64)
				if !ok {
					return models.Set{}
				}
				if float32(weight) != nextSet.Weight {
					nextSet.FactWeight = float32(weight)
				} else {
					nextSet.FactWeight = float32(0)
				}
				return nextSet
			},
			"❌ Неверный формат веса. Введите число (например: 42.5)",
			"✅ Вес обновлен",
		)

	case strings.HasPrefix(state, "awaiting_minutes_"):
		err = s.awaitingEnterData(
			chatID, state,
			func() (interface{}, error) { return strconv.ParseInt(text, 10, 64) },
			func(nextSet models.Set, value interface{}) models.Set {
				minutes, ok := value.(int64)
				if !ok {
					return models.Set{}
				}
				if int(minutes) != nextSet.Minutes {
					nextSet.FactMinutes = int(minutes)
				} else {
					nextSet.FactMinutes = 0
				}
				return nextSet
			},
			"❌ Неверный формат минут. Введите число (например: 42)",
			"✅ Время обновлено",
		)

	case strings.HasPrefix(state, "awaiting_meters_"):
		err = s.awaitingEnterData(
			chatID, state,
			func() (interface{}, error) { return strconv.ParseInt(text, 10, 64) },
			func(nextSet models.Set, value interface{}) models.Set {
				meters, ok := value.(int64)
				if !ok {
					return models.Set{}
				}
				if int(meters) != nextSet.Meters {
					nextSet.FactMeters = int(meters)
				} else {
					nextSet.FactMeters = 0
				}
				return nextSet
			},
			"❌ Неверный формат минут. Введите число (например: 42)",
			"✅ Дистанция обновлена",
		)

	case strings.HasPrefix(state, "awaiting_program_name_"):
		programID, _ := strconv.ParseInt(strings.TrimPrefix(state, "awaiting_program_name_"), 10, 64)
		program, err := s.programsRepo.Get(programID)
		if err != nil {
			return
		}
		program.Name = text
		err = s.programsRepo.Save(&program)
		if err != nil {
			return
		}
		s.programManagement(chatID)

	case strings.HasPrefix(state, "awaiting_day_preset_"):

		text = strings.ToLower(text)

		// parse dayTypeID and exerciseTypeID
		parts := strings.Split(strings.TrimPrefix(state, "awaiting_day_preset_"), "_")
		if len(parts) < 2 {
			return
		}
		dayTypeID, _ := strconv.ParseInt(parts[0], 10, 64)
		exerciseTypeID, _ := strconv.ParseInt(parts[1], 10, 64)
		exerciseType, _ := s.exerciseTypesRepo.Get(exerciseTypeID)

		textArr := strings.Split(text, ":")
		if len(textArr) != 2 {
			s.sendIncorrectPresetMsg(chatID, exerciseType.Units)
			return
		}

		preset := textArr[1]

		units, valid := utils.SplitUnits(textArr[0])
		if !valid {
			s.sendIncorrectPresetMsg(chatID, exerciseType.Units)
			return
		}
		exUnits, _ := utils.SplitUnits(exerciseType.Units)

		if !utils.EqualArrays(exUnits, units) {
			s.sendIncorrectPresetMsg(chatID, exerciseType.Units)
			return
		}
		presetSetLen := 1
		if strings.Contains(preset, "*") {
			presetSetLen = 2
		}
		if len(exUnits) != presetSetLen {
			s.sendIncorrectPresetMsg(chatID, exerciseType.Units)
			return
		}

		if !utils.IsValidPreset(preset) {
			s.sendIncorrectPresetMsg(chatID, exerciseType.Units)
			return
		}

		var dayType models.WorkoutDayType
		dayType, err = s.dayTypesRepo.Get(dayTypeID)
		if err != nil {
			return
		}
		if dayType.Preset != "" {
			dayType.Preset += ";"
		}

		dayType.Preset += fmt.Sprintf("%d:[%s]", exerciseTypeID, preset)
		err = s.dayTypesRepo.Save(&dayType)
		if err != nil {
			return
		}
		s.editProgram(chatID, dayType.WorkoutProgramID)

	case strings.HasPrefix(state, "awaiting_day_name_for_program_"):
		programID, _ := strconv.ParseInt(strings.TrimPrefix(state, "awaiting_day_name_for_program_"), 10, 64)

		dayType, createErr := s.dayTypesRepo.Create(&models.WorkoutDayType{
			WorkoutProgramID: programID,
			Name:             text,
			CreatedAt:        time.Now(),
		})
		if createErr != nil {
			return
		}
		s.addNewDayTypeExercise(chatID, dayType.ID)
	}

	handleErr(method, err)
}

func (s *serviceImpl) sendIncorrectPresetMsg(chatID int64, expectedUnits string) {
	msg := tgbotapi.NewMessage(chatID, "❌ Неверный формат !\n\n"+messages.EnterPreset+
		fmt.Sprintf("\n\n<b>Подсказка:</b> для вашего упражнения следует выбрать <b>%s</b> !", expectedUnits))
	msg.ParseMode = constants.HtmlParseMode
	s.bot.Send(msg)
}

func (s *serviceImpl) awaitingEnterData(
	chatID int64,
	state string,
	parseValue func() (interface{}, error),
	handleSet func(s models.Set, result interface{}) models.Set,
	formatMsg, successMsg string,
) error {
	parts := strings.Split(state, "_")
	if len(parts) < 3 {
		return errors.New("incorrect input")
	}
	exerciseID, _ := strconv.ParseInt(parts[2], 10, 64)

	result, err := parseValue()
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, formatMsg)
		_, err = s.bot.Send(msg)
		if err != nil {
			return err
		}
		return nil
	}

	exercise, _ := s.exercisesRepo.Get(exerciseID)
	nextSet := exercise.NextSet()

	if nextSet.ID != 0 {
		nextSet = handleSet(nextSet, result)
		err = s.setsRepo.Save(&nextSet)
		if err != nil {
			return err
		}

		msg := tgbotapi.NewMessage(chatID, successMsg)
		if _, err = s.bot.Send(msg); err != nil {
			return err
		}
	}
	s.userStatesMachine.SetValue(chatID, "")
	s.showCurrentExerciseSession(chatID, exercise.WorkoutDayID)
	return nil
}

func (s *serviceImpl) export(chatID int64, user *models.User) {
	method := "export"

	groupCodes, _ := s.exerciseGroupTypesRepo.GetAll()
	groupCodesMap := make(map[string]string)
	for _, code := range groupCodes {
		groupCodesMap[code.Code] = code.Name
	}

	workouts, _ := s.workoutsRepo.FindAll(user.ID)
	totalSummary := s.summaryService.BuildTotal(workouts, groupCodesMap)
	byDateSummary := s.summaryService.BuildByDate(workouts)

	exercises, err := s.exercisesRepo.FindAllByUserID(user.ID)
	if err != nil {
		fmt.Printf("%s: error: %v", method, err)
		return
	}

	progresses := make(map[string]map[string]*summarysvc.Progress)
	for _, e := range exercises {
		progresses[e.ExerciseType.Name] = s.summaryService.BuildExerciseProgress(workouts, e.ExerciseType.Name)
	}

	file, err := s.docGeneratorService.ExportToFile(workouts, totalSummary, byDateSummary, progresses, groupCodesMap)
	if err != nil {
		fmt.Println("cannot export file:", err.Error())
		return
	}

	buf, _ := file.WriteToBuffer()
	doc := tgbotapi.FileBytes{Name: "workouts.xlsx", Bytes: buf.Bytes()}

	msg := tgbotapi.NewDocument(chatID, doc)
	_, _ = tghelpers.SendMessage(s.bot, msg, method)
}
