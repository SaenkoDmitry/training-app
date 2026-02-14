package models

import (
	"cmp"
	"fmt"
	"math"
	"strings"
	"time"
)

type Set struct {
	ID int64 `gorm:"primaryKey;autoIncrement"`

	ExerciseID int64
	Exercise   *Exercise `gorm:"foreignKey:ExerciseID;references:ID"` // join

	Reps        int
	FactReps    int
	Weight      float32
	FactWeight  float32
	Minutes     int
	FactMinutes int
	Meters      int
	FactMeters  int
	Completed   bool
	CompletedAt *time.Time
	Index       int
}

func (*Set) TableName() string {
	return "sets"
}

func (s *Set) String(done bool) string {
	var text strings.Builder

	if s.Completed {
		if s.CompletedAt != nil {
			text.WriteString(fmt.Sprintf("✅ [%s]: ", s.CompletedAt.Add(3*time.Hour).Format("15:04:05")))
		}
	} else {
		if done {
			text.WriteString("💔 ")
		} else {
			text.WriteString("🚀 ")
		}
	}
	if s.Exercise.ExerciseType.ContainsMeters() {
		text.WriteString(fmt.Sprintf("%s метров", s.FormatMeters()))
	}
	if s.Exercise.ExerciseType.ContainsMinutes() {
		text.WriteString(fmt.Sprintf("%s минут", s.FormatMinutes()))
	}
	if s.Exercise.ExerciseType.ContainsReps() && s.Exercise.ExerciseType.ContainsWeight() {
		text.WriteString(fmt.Sprintf("%s повт. * %s кг", s.FormatReps(), s.FormatWeight()))
	} else if s.Exercise.ExerciseType.ContainsReps() {
		text.WriteString(fmt.Sprintf("%s повт.", s.FormatReps()))
	}
	if s.Completed {
		//text.WriteString("</strike>")
	}

	text.WriteString("\n")
	return text.String()
}

func strikePlanned[T cmp.Ordered](planned, actual T, completed bool) string {
	if !completed {
		return fmt.Sprintf("%v", planned)
	}
	if planned == actual {
		return fmt.Sprintf("%v", actual)
	}
	return fmt.Sprintf("<strike>%v</strike> <b>%v</b>", planned, actual)
}

func (s *Set) FormatReps() string {
	return strikePlanned(s.Reps, s.FactReps, s.Completed)
}

func formatWeight(weight float32) string {
	// Проверяем, есть ли дробная часть
	if math.Mod(float64(weight), 1) == 0 {
		return fmt.Sprintf("%.0f", weight) // целое число → 0 знаков
	} else {
		return fmt.Sprintf("%.1f", weight) // дробное → 1 знак
	}
}

func (s *Set) FormatWeight() string {
	return strikePlanned(formatWeight(s.Weight), formatWeight(s.FactWeight), s.Completed)
}

func (s *Set) FormatMinutes() string {
	return strikePlanned(s.Minutes, s.FactMinutes, s.Completed)
}

func (s *Set) FormatMeters() string {
	return strikePlanned(s.Meters, s.FactMeters, s.Completed)
}

func (s *Set) GetRealReps() int {
	if s == nil {
		return 0
	}
	if s.FactReps > 0 {
		return s.FactReps
	}
	return s.Reps
}

func (s *Set) GetRealWeight() float32 {
	if s == nil {
		return 0
	}
	if s.FactWeight > 0 {
		return s.FactWeight
	}
	return s.Weight
}

func (s *Set) GetRealMinutes() int {
	if s == nil {
		return 0
	}
	if s.FactMinutes > 0 {
		return s.FactMinutes
	}
	return s.Minutes
}

func (s *Set) GetRealMeters() int {
	if s == nil {
		return 0
	}
	if s.FactMeters > 0 {
		return s.FactMeters
	}
	return s.Meters
}
