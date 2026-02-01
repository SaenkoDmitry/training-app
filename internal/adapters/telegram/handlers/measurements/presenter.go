package measurements

import (
	"fmt"
	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/constants"
	"github.com/SaenkoDmitry/training-tg-bot/internal/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type Presenter struct {
	bot *tgbotapi.BotAPI
}

func NewPresenter(bot *tgbotapi.BotAPI) *Presenter {
	return &Presenter{bot: bot}
}

func (p Presenter) showMenu(chatID int64) {
	msg := tgbotapi.NewMessage(chatID,
		"<b>Выберите действие:</b>")
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➕ Добавить новое", "change_add_new_measurement"),
			tgbotapi.NewInlineKeyboardButtonData("📁 Показать свежие", "measurements_show_top_4_0"),
		),
		//tgbotapi.NewInlineKeyboardRow(
		//	tgbotapi.NewInlineKeyboardButtonData("🗑 Удалить последнее добавленное", "measurements_delete_last"),
		//),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(messages.Export, "export_measurements_to_excel"),
		),
	)
	msg.ParseMode = constants.HtmlParseMode
	msg.ReplyMarkup = keyboard
	p.bot.Send(msg)
}

func (p Presenter) showAllLimitOffset(chatID int64, measurementObjs []dto.Measurement) {
	shoulders := make([]string, 0, len(measurementObjs))
	chests := make([]string, 0, len(measurementObjs))
	hands := make([]string, 0, len(measurementObjs))
	waists := make([]string, 0, len(measurementObjs))
	buttocks := make([]string, 0, len(measurementObjs))
	hips := make([]string, 0, len(measurementObjs))
	calves := make([]string, 0, len(measurementObjs))
	weights := make([]string, 0, len(measurementObjs))

	var from, to string
	if len(measurementObjs) > 0 {
		from = measurementObjs[0].CreatedAt
		to = measurementObjs[len(measurementObjs)-1].CreatedAt
	}
	for _, m := range measurementObjs {
		shoulders = append(shoulders, m.Shoulders)
		chests = append(chests, m.Chest)
		hands = append(hands, m.Hands)
		waists = append(waists, m.Waist)
		buttocks = append(buttocks, m.Buttocks)
		hips = append(hips, m.Hips)
		calves = append(calves, m.Calves)
		weights = append(weights, m.Weight)
	}
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf(
		"<b>%s за период %s – %s</b>\n\n"+
			"• <u>Плечи (см)</u>: %s\n\n"+
			"• <u>Грудь (см)</u>: %s\n\n"+
			"• <u>Руки (см)</u>: %s\n\n"+
			"• <u>Талия (см)</u>: %s\n\n"+
			"• <u>Ягодицы (см)</u>: %s\n\n"+
			"• <u>Бедра (см)</u>: %s\n\n"+
			"• <u>Икры (см)</u>: %s\n\n"+
			"• <u>Вес (кг)</u>: %s",
		messages.Measurements,
		from, to,
		strings.Join(shoulders, delimiter),
		strings.Join(chests, delimiter),
		strings.Join(hands, delimiter),
		strings.Join(waists, delimiter),
		strings.Join(buttocks, delimiter),
		strings.Join(hips, delimiter),
		strings.Join(calves, delimiter),
		strings.Join(weights, delimiter),
	))
	buttons := make([][]tgbotapi.InlineKeyboardButton, 0)
	buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(messages.BackTo, "measurements_menu"),
	))

	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons...)
	msg.ParseMode = constants.HtmlParseMode
	msg.ReplyMarkup = keyboard
	p.bot.Send(msg)
}

const (
	delimiter = " -> "
)
