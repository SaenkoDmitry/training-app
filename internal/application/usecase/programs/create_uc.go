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

func (uc *CreateUseCase) Execute(userID int64, name string) error {
	programObjs, err := uc.programsRepo.FindAll(userID)
	if err != nil {
		return err
	}

	if name == "" {
		name = fmt.Sprintf("#%d", len(programObjs)+1)
	}

	program, err := uc.programsRepo.Create(userID, name)
	if err != nil {
		return err
	}

	if len(programObjs) == 0 {
		user, getUserErr := uc.usersRepo.GetByID(userID)
		if getUserErr != nil {
			return getUserErr
		}

		user.ActiveProgramID = &program.ID
		err = uc.usersRepo.Save(user)
		if err != nil {
			return err
		}
	}

	return nil
}
