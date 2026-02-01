package changes

import (
	"fmt"
	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/constants"
	"github.com/SaenkoDmitry/training-tg-bot/internal/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Presenter struct {
	bot *tgbotapi.BotAPI
}

func NewPresenter(bot *tgbotapi.BotAPI) *Presenter {
	return &Presenter{bot: bot}
}

func (p Presenter) showCreated(chatID int64, createdMeasurement *dto.Measurement) {
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("<b>📅 Дата: %s</b>\n\n"+
		"• <u>Плечи</u>: %s см\n\n"+
		"• <u>Грудь</u>: %s см\n\n"+
		"• <u>Руки</u>: %s см\n\n"+
		"• <u>Талия</u>: %s см\n\n"+
		"• <u>Ягодицы</u>: %s см\n\n"+
		"• <u>Бедра</u>: %s см\n\n"+
		"• <u>Икры</u>: %s см\n\n"+
		"• <u>Вес</u>: %s кг",
		createdMeasurement.CreatedAt,
		createdMeasurement.Shoulders,
		createdMeasurement.Chest,
		createdMeasurement.Hands,
		createdMeasurement.Waist,
		createdMeasurement.Buttocks,
		createdMeasurement.Hips,
		createdMeasurement.Calves,
		createdMeasurement.Weight,
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
