package exercises

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/exercisegrouptypes"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/exercisetypes"
)

type FindTypesByGroupUseCase struct {
	exerciseTypesRepo      exercisetypes.Repo
	exerciseGroupTypesRepo exercisegrouptypes.Repo
}

func NewFindTypesByGroupUseCase(exerciseTypesRepo exercisetypes.Repo, exerciseGroupTypesRepo exercisegrouptypes.Repo) *FindTypesByGroupUseCase {
	return &FindTypesByGroupUseCase{
		exerciseTypesRepo:      exerciseTypesRepo,
		exerciseGroupTypesRepo: exerciseGroupTypesRepo,
	}
}

func (uc *FindTypesByGroupUseCase) Name() string {
	return "Найти упражнения по группе"
}

func (uc *FindTypesByGroupUseCase) Execute(exerciseGroupCode string) (*dto.FindTypesByGroup, error) {
	exerciseTypes, err := uc.exerciseTypesRepo.GetAllByGroup(exerciseGroupCode)
	if err != nil {
		return nil, err
	}

	groups, err := uc.exerciseGroupTypesRepo.GetAll()
	if err != nil {
		return nil, err
	}

	groupsMap := make(map[string]string)
	for _, v := range groups {
		groupsMap[v.Code] = v.Name
	}

	return &dto.FindTypesByGroup{
		ExerciseTypes: mapExerciseTypeDTO(exerciseTypes, groupsMap),
	}, nil
}

func mapExerciseTypeDTO(types []models.ExerciseType, groupsMap map[string]string) []dto.ExerciseTypeDTO {
	result := make([]dto.ExerciseTypeDTO, 0, len(types))
	for _, t := range types {
		result = append(result, dto.ExerciseTypeDTO{
			ID:            t.ID,
			Name:          t.Name,
			Url:           t.Url,
			GroupName:     groupsMap[t.ExerciseGroupTypeCode],
			RestInSeconds: t.RestInSeconds,
			Accent:        t.Accent,
			Units:         t.Units,
			Description:   t.Description,
		})
	}
	return result
}
