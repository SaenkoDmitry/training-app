package daytypes

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/daytypes"
)

type UpdatePresetUseCase struct {
	dayTypesRepo daytypes.Repo
}

func NewUpdatePresetUseCase(
	dayTypesRepo daytypes.Repo,
) *UpdatePresetUseCase {
	return &UpdatePresetUseCase{
		dayTypesRepo: dayTypesRepo,
	}
}

func (uc *UpdatePresetUseCase) Name() string {
	return "Обновить пресет для дня"
}

func (uc *UpdatePresetUseCase) Execute(dayTypeID int64, preset string) error {
	d, err := uc.dayTypesRepo.Get(dayTypeID)
	if err != nil {
		return err
	}
	d.Preset = preset
	err = uc.dayTypesRepo.Save(&d)
	if err != nil {
		return err
	}
	return nil
}
