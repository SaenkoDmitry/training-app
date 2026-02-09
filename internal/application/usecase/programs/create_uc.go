package programs

import (
	"fmt"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/programs"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/users"
)

type CreateUseCase struct {
	programsRepo programs.Repo
	usersRepo    users.Repo
}

func NewCreateUseCase(
	programsRepo programs.Repo,
	usersRepo users.Repo,
) *CreateUseCase {
	return &CreateUseCase{
		programsRepo: programsRepo,
		usersRepo:    usersRepo,
	}
}

func (uc *CreateUseCase) Name() string {
	return "Создать программу"
}

func (uc *CreateUseCase) Execute(chatID int64, name string) error {
	user, err := uc.usersRepo.GetByChatID(chatID)
	if err != nil {
		return err
	}

	programObjs, err := uc.programsRepo.FindAll(user.ID)
	if err != nil {
		return err
	}

	if name == "" {
		name = fmt.Sprintf("#%d", len(programObjs)+1)
	}

	_, err = uc.programsRepo.Create(user.ID, name)
	if err != nil {
		return err
	}

	return nil
}
