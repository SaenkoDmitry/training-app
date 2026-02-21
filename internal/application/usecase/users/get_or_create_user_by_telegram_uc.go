package users

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/users"
)

type GetOrCreateUserByTelegramUseCase struct {
	usersRepo users.Repo
}

func NewGetOrCreateUserByTelegramUseCase(usersRepo users.Repo) *GetOrCreateUserByTelegramUseCase {
	return &GetOrCreateUserByTelegramUseCase{
		usersRepo: usersRepo,
	}
}

func (uc *GetOrCreateUserByTelegramUseCase) Name() string {
	return "Создать телеграм аккаунт в системе или найти существующий"
}

func (uc *GetOrCreateUserByTelegramUseCase) Execute(tgUser dto.TelegramUser) (*models.User, error) {
	user, err := uc.usersRepo.GetByChatID(tgUser.ID)
	if err != nil && !errors.Is(err, users.NotFoundUserErr) {
		return nil, err
	}

	if err != nil && errors.Is(err, users.NotFoundUserErr) {
		user, err = uc.usersRepo.CreateTelegram(&tgbotapi.User{
			ID:           tgUser.ID,
			FirstName:    tgUser.FirstName,
			LastName:     tgUser.LastName,
			UserName:     tgUser.Username,
			LanguageCode: tgUser.LanguageCode,
		})
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
