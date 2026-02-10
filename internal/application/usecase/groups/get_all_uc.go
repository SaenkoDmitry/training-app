package groups

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/exercisegrouptypes"
)

type GetAllUseCase struct {
	exerciseGroupTypesRepo exercisegrouptypes.Repo
}

func NewGetAllUseCase(exerciseGroupTypesRepo exercisegrouptypes.Repo) *GetAllUseCase {
	return &GetAllUseCase{
		exerciseGroupTypesRepo: exerciseGroupTypesRepo,
	}
}

func (uc *GetAllUseCase) Name() string {
	return "Список групп упражнений"
}

func (uc *GetAllUseCase) Execute() (*dto.ExerciseGroupTypeList, error) {
	groups, err := uc.exerciseGroupTypesRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return &dto.ExerciseGroupTypeList{
		Groups: mapGroupDTO(groups),
	}, nil
}

func mapGroupDTO(groups []models.ExerciseGroupType) []dto.Group {
	result := make([]dto.Group, 0, len(groups))
	for _, g := range groups {
		result = append(result, dto.Group{
			Code: g.Code,
			Name: g.Name,
		})
	}
	return result
}
