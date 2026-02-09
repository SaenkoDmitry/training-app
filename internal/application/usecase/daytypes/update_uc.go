package daytypes

import (
	"fmt"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/daytypes"
)

type UpdateUseCase struct {
	dayTypesRepo daytypes.Repo
}

func NewUpdateUseCase(
	dayTypesRepo daytypes.Repo,
) *UpdateUseCase {
	return &UpdateUseCase{
		dayTypesRepo: dayTypesRepo,
	}
}

func (uc *UpdateUseCase) Name() string {
	return "Обновить день"
}

func (uc *UpdateUseCase) Execute(dayTypeID, exerciseTypeID int64, preset string) error {
	d, err := uc.dayTypesRepo.Get(dayTypeID)
	if err != nil {
		return err
	}
	d.Preset += fmt.Sprintf(";%d:[%s]", exerciseTypeID, preset)
	err = uc.dayTypesRepo.Save(&d)
	if err != nil {
		return err
	}
	return nil
}
