package dto

import "github.com/SaenkoDmitry/training-tg-bot/internal/models"

type CurrentExerciseSession struct {
	Exercise      models.Exercise
	ExerciseObj   models.ExerciseType
	DayType       models.WorkoutDayType
	WorkoutDay    models.WorkoutDay
	ExerciseIndex int
}

type ExerciseTypeList struct {
	ExerciseTypes []models.ExerciseType
}

type FindTypesByGroup struct {
	ExerciseTypes []ExerciseTypeDTO `json:"exercise_types"`
}

type ExerciseTypeDTO struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Url           string `json:"url"`
	GroupName     string `json:"group_name"`
	RestInSeconds int    `json:"rest_in_seconds"`
	Accent        string `json:"accent"`
	Units         string `json:"units"`
	Description   string `json:"description"`
}

type ConfirmDeleteExercise struct {
	Exercise    models.Exercise
	ExerciseObj models.ExerciseType
}

type GetExercise struct {
	ExerciseType models.ExerciseType
}

type CreateExercise struct {
	ExerciseObj models.ExerciseType
}
