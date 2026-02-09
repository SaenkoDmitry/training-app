package daytypes

import (
	"bytes"
	"fmt"
	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/daytypes"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/exercisetypes"
	"github.com/SaenkoDmitry/training-tg-bot/internal/utils"
	"time"
)

type GetUseCase struct {
	dayTypesRepo      daytypes.Repo
	exerciseTypesRepo exercisetypes.Repo
}

func NewGetUseCase(
	dayTypesRepo daytypes.Repo,
	exerciseTypesRepo exercisetypes.Repo,
) *GetUseCase {
	return &GetUseCase{
		dayTypesRepo:      dayTypesRepo,
		exerciseTypesRepo: exerciseTypesRepo,
	}
}

func (uc *GetUseCase) Name() string {
	return "Загрузить тренировочный день"
}

func (uc *GetUseCase) Execute(dayTypeID int64) (*dto.WorkoutDayTypeDTO, error) {
	dayType, err := uc.dayTypesRepo.Get(dayTypeID)
	if err != nil {
		return nil, err
	}

	exerciseTypesMap := make(map[int64]models.ExerciseType)
	exTypes, err := uc.exerciseTypesRepo.GetAll()
	if err != nil {
		return nil, err
	}
	for _, ex := range exTypes {
		exerciseTypesMap[ex.ID] = ex
	}

	return &dto.WorkoutDayTypeDTO{
		ID:               dayType.ID,
		WorkoutProgramID: dayType.WorkoutProgramID,
		Name:             dayType.Name,
		Preset:           formatPreset(dayType.Preset, exerciseTypesMap),
		CreatedAt:        "📅 " + dayType.CreatedAt.Add(time.Hour*3).Format("02.01.2006 15:04"),
	}, nil
}

func formatPreset(preset string, exerciseTypesMap map[int64]models.ExerciseType) string {
	exercises := utils.SplitPreset(preset)
	buffer := &bytes.Buffer{}
	for i, ex := range exercises {
		exerciseType, ok := exerciseTypesMap[ex.ID]
		if !ok {
			continue
		}
		buffer.WriteString(fmt.Sprintf("• <b>%d.</b> <u>%s</u>\n", i+1, exerciseType.Name))
		buffer.WriteString(fmt.Sprintf("    • "))
		for i, set := range ex.Sets {
			if i > 0 {
				buffer.WriteString(", ")
			}
			if set.Minutes > 0 {
				buffer.WriteString(fmt.Sprintf("%d мин", set.Minutes))
			} else {
				buffer.WriteString(fmt.Sprintf("%d * %.0f кг", set.Reps, set.Weight))
			}
		}
		buffer.WriteString("\n\n")
	}
	return buffer.String()
}
