package docgenerator

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
	summarysvc "github.com/SaenkoDmitry/training-tg-bot/internal/service/summary"
	"github.com/SaenkoDmitry/training-tg-bot/internal/utils"
	"github.com/xuri/excelize/v2"
)

type Service interface {
	ExportWorkoutsToFile(
		workouts []models.WorkoutDay,
		summary map[string]*summarysvc.ExerciseSummary,
		byDateSummary map[string]*summarysvc.DateSummary,
		exerciseProgressByDates []*summarysvc.ExerciseProgressByDates,
		groupCodesMap map[string]string,
		typeSummary map[utils.DateRange]map[string]*summarysvc.WeekSummary,
	) (*excelize.File, error)

	ExportMeasurementsToFile(measurements []*dto.Measurement) (*excelize.File, error)
}

type serviceImpl struct {
	summaryService summarysvc.Service
}

func NewService(summaryService summarysvc.Service) Service {
	return &serviceImpl{
		summaryService: summaryService,
	}
}
