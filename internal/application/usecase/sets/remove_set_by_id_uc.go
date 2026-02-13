package sets

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/sets"
)

type RemoveByIDUseCase struct {
	setsRepo sets.Repo
}

func NewRemoveByIDUseCase(
	setsRepo sets.Repo,
) *RemoveByIDUseCase {
	return &RemoveByIDUseCase{
		setsRepo: setsRepo,
	}
}

func (uc *RemoveByIDUseCase) Name() string {
	return "Удалить подход"
}

func (uc *RemoveByIDUseCase) Execute(setID int64) error {
	err := uc.setsRepo.Delete(setID)
	if err != nil {
		return err
	}

	return nil
}
