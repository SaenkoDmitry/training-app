package programs

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/exercisetypes"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/programs"
)

type GetUseCase struct {
	programsRepo     programs.Repo
	exerciseTypeRepo exercisetypes.Repo
}

func NewGetUseCase(
	programsRepo programs.Repo,
	exerciseTypeRepo exercisetypes.Repo,
) *GetUseCase {
	return &GetUseCase{
		programsRepo:     programsRepo,
		exerciseTypeRepo: exerciseTypeRepo,
	}
}

func (uc *GetUseCase) Name() string {
	return "Редактировать программу"
}

func (uc *GetUseCase) Execute(programID int64) (*dto.GetProgramDTO, error) {
	program, err := uc.programsRepo.Get(programID)
	if err != nil {
		return nil, err
	}

	return &dto.GetProgramDTO{
		Program: program,
	}, nil
}
