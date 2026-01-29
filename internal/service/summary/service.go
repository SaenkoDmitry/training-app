package summary

import (
	"math"
	"time"

	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
)

type Service interface {
	BuildTotal(workouts []models.WorkoutDay, groupCodesMap map[string]string) map[string]*ExerciseSummary
	BuildByDate(workouts []models.WorkoutDay) map[string]*DateSummary
	BuildExerciseProgress(workouts []models.WorkoutDay, exerciseName string) map[string]*Progress
}

type serviceImpl struct {
}

func NewService() Service {
	return &serviceImpl{}
}

func (s *serviceImpl) BuildTotal(workouts []models.WorkoutDay, groupCodesMap map[string]string) map[string]*ExerciseSummary {
	summary := make(map[string]*ExerciseSummary)

	for _, w := range workouts {
		if !w.Completed {
			continue
		}
		date := w.StartedAt.Add(3 * time.Hour).Format("2006-01-02")

		for _, e := range w.Exercises {
			if e.CompletedSets() == 0 {
				continue
			}

			sum, ok := summary[e.ExerciseType.Name]
			if !ok {
				sum = &ExerciseSummary{
					Workouts: make(map[string]struct{}),
				}
				summary[e.ExerciseType.Name] = sum
			}

			sum.ExerciseType = groupCodesMap[e.ExerciseType.ExerciseGroupTypeCode]
			sum.Workouts[date] = struct{}{}

			for _, set := range e.Sets {
				if !set.Completed {
					continue
				}
				sum.Sets++
				sum.TotalWeight += float64(set.Weight) * float64(set.Reps)
				sum.TotalReps += set.Reps
				sum.TotalMinutes += set.Minutes

				if set.Weight > sum.MaxWeight {
					sum.MaxWeight = set.Weight
				}
			}
		}
	}

	for _, sum := range summary {
		if sum.TotalReps > 0 {
			sum.AvgWeight = math.Round(sum.TotalWeight / float64(sum.TotalReps))
		}
	}

	return summary
}

func (s *serviceImpl) BuildByDate(workouts []models.WorkoutDay) map[string]*DateSummary {
	result := make(map[string]*DateSummary)

	for _, w := range workouts {
		date := w.StartedAt.Add(3 * time.Hour).Format("2006-01-02")

		d, ok := result[date]
		if !ok {
			d = &DateSummary{
				Workouts:  1,
				Exercises: make(map[string]struct{}),
			}
			result[date] = d
		}

		for _, e := range w.Exercises {
			d.Exercises[e.ExerciseType.Name] = struct{}{}

			for _, sum := range e.Sets {
				d.Sets++
				d.TotalVolume += sum.Weight * float32(sum.Reps)

				if sum.Weight > d.MaxWeight {
					d.MaxWeight = sum.Weight
				}
			}
		}
	}

	return result
}

func (s *serviceImpl) BuildExerciseProgress(
	workouts []models.WorkoutDay,
	exerciseName string,
) map[string]*Progress {

	progress := make(map[string]*Progress)

	for _, w := range workouts {
		if !w.Completed {
			continue
		}

		date := w.StartedAt.Add(3 * time.Hour).Format("2006-01-02")

		for _, e := range w.Exercises {
			if e.ExerciseType.Name != exerciseName {
				continue
			}
			if e.CompletedSets() == 0 {
				continue
			}

			sumWeight := float32(0)
			countOfReps := 0

			progress[date] = &Progress{}

			for _, set := range e.Sets {
				if !set.Completed {
					continue
				}

				countOfReps += set.GetRealReps()
				sumWeight += set.GetRealWeight() * float32(set.GetRealReps())
				if progress[date].MaxWeight < set.GetRealWeight() ||
					progress[date].MaxWeight == set.GetRealWeight() && progress[date].MaxReps < set.GetRealReps() {
					progress[date].MaxWeight = set.GetRealWeight()
					progress[date].MaxReps = set.GetRealReps()
				}

				progress[date].SumMinutes += set.GetRealMinutes()
				if progress[date].MinMinutes == 0 {
					progress[date].MinMinutes = set.GetRealMinutes()
					progress[date].MaxMinutes = set.GetRealMinutes()
				} else {
					progress[date].MinMinutes = min(progress[date].MinMinutes, set.GetRealMinutes())
					progress[date].MaxMinutes = max(progress[date].MaxMinutes, set.GetRealMinutes())
				}

				progress[date].SumMeters += set.GetRealMeters()
				if progress[date].MinMeters == 0 {
					progress[date].MinMeters = set.GetRealMeters()
					progress[date].MaxMeters = set.GetRealMeters()
				} else {
					progress[date].MinMeters = min(progress[date].MinMeters, set.GetRealMeters())
					progress[date].MaxMeters = max(progress[date].MaxMeters, set.GetRealMeters())
				}
			}
			progress[date].AvgWeight = sumWeight / float32(countOfReps)
		}
	}

	return progress
}
