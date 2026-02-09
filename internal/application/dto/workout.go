package dto

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
)

type WorkoutItem struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	StartedAt string `json:"started_at"`
	Duration  string `json:"duration"`
	Completed bool   `json:"completed"`
	Status    string `json:"status"`
}

type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}

type ShowMyWorkoutsResult struct {
	Items      []WorkoutItem `json:"items"`
	Pagination Pagination    `json:"pagination"`
}

type ConfirmDeleteWorkout struct {
	WorkoutID   int64
	DayTypeName string
}

type DeleteWorkout struct {
}

type ConfirmFinishWorkout struct {
	DayType models.WorkoutDayType
}

type FinishWorkout struct {
	WorkoutID int64
}

type CreateWorkout struct {
	WorkoutID int64
}

type StartWorkout struct {
}

type FormattedWorkout struct {
	ID          int64                `json:"id"`
	UserID      int64                `json:"user_id"`
	Status      string               `json:"status"`
	StartedAt   string               `json:"started_at"`
	Duration    string               `json:"duration"`
	EndedAt     string               `json:"ended_at"`
	DayTypeName string               `json:"day_type_name"`
	Completed   bool                 `json:"completed"`
	Exercises   []*FormattedExercise `json:"exercises"`
}

type FormattedExercise struct {
	ID            int64           `json:"id"`
	Name          string          `json:"name"`
	Url           string          `json:"url"`
	GroupName     string          `json:"group_name"`
	RestInSeconds int             `json:"rest_in_seconds"`
	Accent        string          `json:"accent"`
	Units         string          `json:"units"`
	Description   string          `json:"description"`
	Index         int             `json:"index"`
	Sets          []*FormattedSet `json:"sets"`
	SumWeight     float32         `json:"sum_weight"`
}

type FormattedSet struct {
	ID              int64  `json:"id"`
	FormattedString string `json:"formatted_string"`
	Completed       bool   `json:"completed"`
	CompletedAt     string `json:"completed_at"`
	Index           int    `json:"index"`
}

type WorkoutProgress struct {
	Workout *FormattedWorkout `json:"workout"`

	TotalExercises     int
	CompletedExercises int

	TotalSets     int
	CompletedSets int

	ProgressPercent int
	RemainingMin    *int

	SessionStarted bool
}

type WorkoutStatistic struct {
	DayType            models.WorkoutDayType
	WorkoutDay         models.WorkoutDay
	TotalWeight        float64
	CompletedExercises int
	CardioTime         int
	ExerciseTypesMap   map[int64]models.ExerciseType
	ExerciseWeightMap  map[int64]float64
	ExerciseTimeMap    map[int64]int
}

type ShowWorkoutByUserID struct {
	Workouts []models.WorkoutDay
	User     *models.User
}
