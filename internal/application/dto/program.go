package dto

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
)

type DeleteProgramResult struct {
	User *models.User
}

type ActivateProgramResult struct {
	User *models.User
}

type GetAllPrograms struct {
	User     *models.User
	Programs []*ProgramDTO `json:"programs"`
}

type ProgramDTO struct {
	ID        int64                `json:"id"`
	UserID    int64                `json:"user_id"`
	Name      string               `json:"name"`
	CreatedAt string               `json:"created_at"`
	DayTypes  []*WorkoutDayTypeDTO `json:"day_types"`
}

type WorkoutDayTypeDTO struct {
	ID               int64  `json:"id"`
	WorkoutProgramID int64  `json:"program_id"`
	Name             string `json:"name"`
	Preset           string `json:"preset"`
	CreatedAt        string `json:"created_at"`
}

type CreateProgramResult struct {
	User *models.User
}

type ListGroups struct {
	Groups []models.ExerciseGroupType
}

type GetProgramDTO struct {
	Program models.WorkoutProgram
}

type ConfirmDeleteProgram struct {
	Program models.WorkoutProgram
}

type RenameProgram struct {
	Program models.WorkoutProgram
}
