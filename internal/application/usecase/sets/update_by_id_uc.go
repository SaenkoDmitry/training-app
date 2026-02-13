package sets

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/sets"
)

type UpdateByIDUseCase struct {
	setsRepo sets.Repo
}

func NewUpdateByIDUseCase(
	setsRepo sets.Repo,
) *UpdateByIDUseCase {
	return &UpdateByIDUseCase{
		setsRepo: setsRepo,
	}
}

func (uc *UpdateByIDUseCase) Name() string {
	return "Изменить подход"
}

func (uc *UpdateByIDUseCase) Execute(setID int64, newSetDTO *dto.NewSet) error {
	set, err := uc.setsRepo.Get(setID)
	if err != nil {
		return err
	}

	set.FactReps = int(newSetDTO.NewReps)
	set.FactWeight = float32(newSetDTO.NewWeight)
	set.FactMinutes = int(newSetDTO.NewMinutes)
	set.FactMeters = int(newSetDTO.NewMeters)

	err = uc.setsRepo.Save(set)
	if err != nil {
		return err
	}

	return nil
}
