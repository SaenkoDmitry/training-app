package measurements

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/measurements"
	"github.com/SaenkoDmitry/training-tg-bot/internal/utils"
)

type CreateUseCase struct {
	measurementsRepo measurements.Repo
}

func NewCreateUseCase(measurementsRepo measurements.Repo) *CreateUseCase {
	return &CreateUseCase{
		measurementsRepo: measurementsRepo,
	}
}

func (uc *CreateUseCase) Name() string {
	return "Добавить новый измерение тела"
}

func (uc *CreateUseCase) Execute(measurement *models.Measurement) (*dto.Measurement, error) {
	err := uc.measurementsRepo.Save(measurement)
	if err != nil {
		return nil, err
	}

	return &dto.Measurement{
		CreatedAt: utils.FormatDate(measurement.CreatedAt),
		Shoulders: utils.FormatCentimeters(measurement.Shoulders),
		Chest:     utils.FormatCentimeters(measurement.Chest),
		Hands:     utils.FormatCentimeters(measurement.Hands),
		Waist:     utils.FormatCentimeters(measurement.Waist),
		Buttocks:  utils.FormatCentimeters(measurement.Buttocks),
		Hips:      utils.FormatCentimeters(measurement.Hips),
		Calves:    utils.FormatCentimeters(measurement.Calves),
		Weight:    utils.FormatKilograms(measurement.Weight),
	}, nil
}
