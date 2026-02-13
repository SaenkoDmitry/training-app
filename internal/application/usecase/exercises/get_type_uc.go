package exercises

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/exercises"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/exercisetypes"
)

type GetTypeUseCase struct {
	exercisesRepo     exercises.Repo
	exerciseTypesRepo exercisetypes.Repo
}

func NewGetTypeUseCase(
	exercisesRepo exercises.Repo,
	exerciseTypesRepo exercisetypes.Repo,
) *GetTypeUseCase {
	return &GetTypeUseCase{
		exercisesRepo:     exercisesRepo,
		exerciseTypesRepo: exerciseTypesRepo,
	}
}

func (uc *GetTypeUseCase) Name() string {
	return "Показать данные о типе упражнения"
}

func (uc *GetTypeUseCase) Execute(exerciseTypeID int64) (*dto.GetExerciseType, error) {
	exType, err := uc.exerciseTypesRepo.Get(exerciseTypeID)
	if err != nil {
		return nil, err
	}

	return &dto.GetExerciseType{ExerciseType: exType}, nil
}
