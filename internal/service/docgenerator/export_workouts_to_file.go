package docgenerator

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/constants"
	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
	"github.com/SaenkoDmitry/training-tg-bot/internal/service/docgenerator/helpers"
	summarysvc "github.com/SaenkoDmitry/training-tg-bot/internal/service/summary"
	"github.com/SaenkoDmitry/training-tg-bot/internal/utils"
	"github.com/xuri/excelize/v2"
)

const (
	DefaultSheet = "Sheet1"

	WorkoutSheet                = "Все тренировки"
	TotalSummarySheet           = "Упражнения"
	ByDateSummarySheet          = "По датам"
	ByWeekAndExTypeSummarySheet = "По неделям & типу упражнения"
	ByExerciseSummarySheet      = "Динамика"
)

func (s *serviceImpl) ExportWorkoutsToFile(
	workouts []models.WorkoutDay,
	summary map[string]*summarysvc.ExerciseSummary,
	byDateSummary map[string]*summarysvc.DateSummary,
	exerciseProgressByDates []*summarysvc.ExerciseProgressByDates,
	groupCodesMap map[string]string,
	byWeekAndExerciseTypeSummary map[utils.DateRange]map[string]*summarysvc.WeekSummary,
) (*excelize.File, error) {
	f := excelize.NewFile()

	redHeaderStyle := helpers.HeaderStyle(f, constants.RedColor)
	greedHeaderStyle := helpers.HeaderStyle(f, constants.GreenColor)
	blueHeaderStyle := helpers.HeaderStyle(f, constants.BlueColor)

	s.writeWorkoutsSheet(f, workouts, groupCodesMap)
	s.writeTotalSummarySheet(f, summary)
	s.writeByDateSummarySheet(f, byDateSummary)
	s.writeByWeekAndExTypeSummarySheet(f, byWeekAndExerciseTypeSummary)
	s.writeWorkoutProgressChartsSheet(f, exerciseProgressByDates, redHeaderStyle, greedHeaderStyle, blueHeaderStyle)

	_ = f.SetRowStyle(WorkoutSheet, 1, 1, blueHeaderStyle)
	_ = f.SetRowStyle(TotalSummarySheet, 1, 1, redHeaderStyle)
	_ = f.SetRowStyle(ByWeekAndExTypeSummarySheet, 1, 1, greedHeaderStyle)
	_ = f.SetRowStyle(ByDateSummarySheet, 1, 1, greedHeaderStyle)

	helpers.AutoFitColumns(f, WorkoutSheet, 1, 8)
	helpers.AutoFitColumns(f, TotalSummarySheet, 1, 7)
	helpers.AutoFitColumns(f, ByWeekAndExTypeSummarySheet, 1, 10)
	helpers.AutoFitColumns(f, ByDateSummarySheet, 1, 6)
	helpers.AutoFitColumns(f, ByExerciseSummarySheet, 1, 4)

	_ = f.DeleteSheet(DefaultSheet)

	f.SetActiveSheet(0)
	return f, nil
}
