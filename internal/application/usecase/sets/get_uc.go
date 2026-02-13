package sets

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/sets"
)

type GetByIDUseCase struct {
	setsRepo sets.Repo
}

func NewGetByIDUseCase(
	setsRepo sets.Repo,
) *GetByIDUseCase {
	return &GetByIDUseCase{
		setsRepo: setsRepo,
	}
}

func (uc *GetByIDUseCase) Name() string {
	return "Получить подход"
}

func (uc *GetByIDUseCase) Execute(setID int64) (*models.Set, error) {
	set, err := uc.setsRepo.Get(setID)
	if err != nil {
		return nil, err
	}

	return set, nil
}
