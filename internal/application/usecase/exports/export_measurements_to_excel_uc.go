package exports

import (
	"bytes"
	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/measurements"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/users"
	"github.com/SaenkoDmitry/training-tg-bot/internal/service/docgenerator"
	"github.com/SaenkoDmitry/training-tg-bot/internal/utils"
	"sort"
)

type ExportMeasurementsToExcelUseCase struct {
	usersRepo           users.Repo
	measurementsRepo    measurements.Repo
	docGeneratorService docgenerator.Service
}

func NewExportMeasurementsToExcelUseCase(
	usersRepo users.Repo,
	measurementsRepo measurements.Repo,
	docGeneratorService docgenerator.Service,
) *ExportMeasurementsToExcelUseCase {
	return &ExportMeasurementsToExcelUseCase{
		usersRepo:           usersRepo,
		measurementsRepo:    measurementsRepo,
		docGeneratorService: docGeneratorService,
	}
}

func (uc *ExportMeasurementsToExcelUseCase) Name() string {
	return "Экспорт в Excel"
}

func (uc *ExportMeasurementsToExcelUseCase) Execute(chatID int64) (*bytes.Buffer, error) {
	user, err := uc.usersRepo.GetByChatID(chatID)
	if err != nil {
		return nil, err
	}

	measurementObjs, err := uc.measurementsRepo.FindAll(user.ID)
	if err != nil {
		return nil, err
	}

	sort.Slice(measurementObjs, func(i, j int) bool {
		return measurementObjs[i].CreatedAt.Before(measurementObjs[j].CreatedAt)
	})

	measurementDTOs := make([]*dto.Measurement, 0)
	for _, m := range measurementObjs {
		measurementDTOs = append(measurementDTOs, &dto.Measurement{
			CreatedAt: utils.FormatDate(m.CreatedAt),
			Shoulders: utils.FormatCentimeters(m.Shoulders),
			Chest:     utils.FormatCentimeters(m.Chest),
			HandLeft:  utils.FormatCentimeters(m.HandLeft),
			HandRight: utils.FormatCentimeters(m.HandRight),
			Waist:     utils.FormatCentimeters(m.Waist),
			Buttocks:  utils.FormatCentimeters(m.Buttocks),
			HipLeft:   utils.FormatCentimeters(m.HipLeft),
			HipRight:  utils.FormatCentimeters(m.HipRight),
			CalfLeft:  utils.FormatCentimeters(m.CalfLeft),
			CalfRight: utils.FormatCentimeters(m.CalfRight),
			Weight:    utils.FormatKilograms(m.Weight),
		})
	}

	file, err := uc.docGeneratorService.ExportMeasurementsToFile(measurementDTOs)
	if err != nil {
		return nil, err
	}

	buf, err := file.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buf, nil
}
