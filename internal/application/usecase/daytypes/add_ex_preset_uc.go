package daytypes

import (
	"fmt"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/daytypes"
)

type AddExPresetUseCase struct {
	dayTypesRepo daytypes.Repo
}

func NewAddExPresetUseCase(
	dayTypesRepo daytypes.Repo,
) *AddExPresetUseCase {
	return &AddExPresetUseCase{
		dayTypesRepo: dayTypesRepo,
	}
}

func (uc *AddExPresetUseCase) Name() string {
	return "Добавить пресет для упражнения"
}

func (uc *AddExPresetUseCase) Execute(dayTypeID, exerciseTypeID int64, preset string) error {
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
