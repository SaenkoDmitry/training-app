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
			tgbotapi.NewInlineKeyboardButtonData("📁 Показать топ-10", "measurements_show_top_10_0"),
		),
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
		"<b>Измерения:</b>\n\n"+
			"• <u>Плечи (см)</u>: %s\n\n"+
			"• <u>Грудь (см)</u>: %s см\n\n"+
			"• <u>Руки (см)</u>: %s см\n\n"+
			"• <u>Талия (см)</u>: %s см\n\n"+
			"• <u>Ягодицы (см)</u>: %s см\n\n"+
			"• <u>Бедра (см)</u>: %s см\n\n"+
			"• <u>Икры (см)</u>: %s см\n\n"+
			"• <u>Вес (кг)</u>: %s",
		strings.Join(shoulders, "->"),
		strings.Join(chests, "->"),
		strings.Join(hands, "->"),
		strings.Join(waists, "->"),
		strings.Join(buttocks, "->"),
		strings.Join(hips, "->"),
		strings.Join(calves, "->"),
		strings.Join(weights, "->"),
	))
	msg.ParseMode = constants.HtmlParseMode
	p.bot.Send(msg)
}
