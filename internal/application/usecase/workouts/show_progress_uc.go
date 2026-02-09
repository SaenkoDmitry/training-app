package workouts

import (
	"errors"
	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/exercisegrouptypes"
	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/sessions"
	"github.com/SaenkoDmitry/training-tg-bot/internal/utils"
	"time"

	"github.com/SaenkoDmitry/training-tg-bot/internal/repository/workouts"
)

type ShowProgressUseCase struct {
	workoutsRepo           workouts.Repo
	sessionsRepo           sessions.Repo
	exerciseGroupTypesRepo exercisegrouptypes.Repo
}

func NewShowProgressUseCase(
	workoutsRepo workouts.Repo,
	sessionsRepo sessions.Repo,
	exerciseGroupTypesRepo exercisegrouptypes.Repo,
) *ShowProgressUseCase {
	return &ShowProgressUseCase{
		workoutsRepo:           workoutsRepo,
		sessionsRepo:           sessionsRepo,
		exerciseGroupTypesRepo: exerciseGroupTypesRepo,
	}
}

func (uc *ShowProgressUseCase) Name() string {
	return "Показать прогресс тренировки"
}

var (
	ErrWorkoutNotFound = errors.New("тренировка не найдена")
)

func (uc *ShowProgressUseCase) Execute(workoutID int64) (*dto.WorkoutProgress, error) {
	w, err := uc.workoutsRepo.Get(workoutID)
	if err != nil || w.ID == 0 {
		return nil, ErrWorkoutNotFound
	}

	totalExercises := len(w.Exercises)
	totalSets := 0
	completedExercises := 0
	completedSets := 0

	for _, exercise := range w.Exercises {
		setsCount := len(exercise.Sets)
		totalSets += setsCount

		done := exercise.CompletedSets()
		completedSets += done

		if done == setsCount && setsCount > 0 {
			completedExercises++
		}
	}

	progress := 0
	if totalSets > 0 {
		progress = (completedSets * 100) / totalSets
	}

	var remaining *int
	if w.EndedAt == nil && completedSets > 0 {
		elapsed := time.Since(w.StartedAt)
		setsPerMinute := float64(completedSets) / elapsed.Minutes()

		if setsPerMinute > 0 {
			left := totalSets - completedSets
			min := int(float64(left) / setsPerMinute)
			remaining = &min
		}
	}

	groups, err := uc.exerciseGroupTypesRepo.GetAll()
	if err != nil {
		return nil, err
	}

	session, _ := uc.sessionsRepo.GetByWorkoutID(workoutID)

	return &dto.WorkoutProgress{
		Workout:            mapToFormattedWorkout(w, groups),
		TotalExercises:     totalExercises,
		CompletedExercises: completedExercises,
		TotalSets:          totalSets,
		CompletedSets:      completedSets,
		ProgressPercent:    progress,
		RemainingMin:       remaining,
		SessionStarted:     session.IsActive,
	}, nil
}

func mapToFormattedWorkout(w models.WorkoutDay, groups []models.ExerciseGroupType) *dto.FormattedWorkout {
	groupsMap := make(map[string]string)
	for _, v := range groups {
		groupsMap[v.Code] = v.Name
	}

	res := &dto.FormattedWorkout{
		ID:          w.ID,
		UserID:      w.UserID,
		StartedAt:   "📆️ " + utils.FormatDateTimeWithDayOfWeek(w.StartedAt),
		Status:      w.Status(),
		Duration:    utils.BetweenTimes(w.StartedAt, w.EndedAt),
		DayTypeName: w.WorkoutDayType.Name,
		Completed:   w.Completed,
	}
	for _, ex := range w.Exercises {
		sumWeight := float32(0)
		sets := make([]*dto.FormattedSet, 0, len(ex.Sets))
		for _, s := range ex.Sets {
			if s.Completed {
				sumWeight += s.GetRealWeight() * float32(s.GetRealReps())
			}
			newSet := &dto.FormattedSet{
				ID:              s.ID,
				FormattedString: s.String(w.Completed),
				Completed:       s.Completed,
				Index:           s.Index,
			}
			if s.CompletedAt != nil {
				newSet.CompletedAt = s.CompletedAt.Add(3 * time.Hour).Format("15:04:05")
			}
			sets = append(sets, newSet)
		}
		res.Exercises = append(res.Exercises, &dto.FormattedExercise{
			ID:            ex.ID,
			Name:          ex.ExerciseType.Name,
			Units:         ex.ExerciseType.Units,
			GroupName:     groupsMap[ex.ExerciseType.ExerciseGroupTypeCode],
			RestInSeconds: ex.ExerciseType.RestInSeconds,
			Accent:        ex.ExerciseType.Accent,
			Description:   ex.ExerciseType.Description,
			SumWeight:     sumWeight,
			Index:         ex.Index,
			Sets:          sets,
		})
	}
	if w.EndedAt != nil {
		res.EndedAt = utils.FormatDate(*w.EndedAt)
	}
	return res
}
