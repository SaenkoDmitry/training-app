package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SaenkoDmitry/training-tg-bot/internal/api/helpers"
	"github.com/SaenkoDmitry/training-tg-bot/internal/application/dto"
	"github.com/SaenkoDmitry/training-tg-bot/internal/middlewares"
)

func (s *serviceImpl) AddSet(w http.ResponseWriter, r *http.Request) {
	claims, ok := middlewares.FromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	exerciseID, err := helpers.ParseInt64Param("exercise_id", w, r)
	if err != nil {
		return
	}
	_, err = s.container.AddOneMoreSetUC.Execute(exerciseID)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	err = s.validateAccessToExercise(w, claims.ChatID, exerciseID)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}

func (s *serviceImpl) DeleteSet(w http.ResponseWriter, r *http.Request) {
	claims, ok := middlewares.FromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	setID, err := helpers.ParseInt64Param("id", w, r)
	if err != nil {
		return
	}

	set, err := s.container.GetSetByIDUC.Execute(setID)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	err = s.validateAccessToExercise(w, claims.ChatID, set.ExerciseID)
	if err != nil {
		return
	}

	err = s.container.RemoveSetByIDUC.Execute(setID)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}

func (s *serviceImpl) CompleteSet(w http.ResponseWriter, r *http.Request) {
	claims, ok := middlewares.FromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	setID, err := helpers.ParseInt64Param("id", w, r)
	if err != nil {
		return
	}

	set, err := s.container.GetSetByIDUC.Execute(setID)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	err = s.validateAccessToExercise(w, claims.ChatID, set.ExerciseID)
	if err != nil {
		return
	}

	err = s.container.CompleteByIDSetUC.Execute(setID)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}

func (s *serviceImpl) ChangeSet(w http.ResponseWriter, r *http.Request) {
	claims, ok := middlewares.FromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	setID, err := helpers.ParseInt64Param("id", w, r)
	if err != nil {
		return
	}

	// Разбираем JSON из тела запроса
	var input struct {
		FactReps    int     `json:"fact_reps"`
		FactWeight  float32 `json:"fact_weight"`
		FactMinutes int     `json:"fact_minutes"`
		FactMeters  int     `json:"fact_meters"`
	}

	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	set, err := s.container.GetSetByIDUC.Execute(setID)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	err = s.validateAccessToExercise(w, claims.ChatID, set.ExerciseID)
	if err != nil {
		return
	}

	err = s.container.UpdateSetByIDUC.Execute(setID, &dto.NewSet{
		NewReps:    int64(input.FactReps),
		NewWeight:  float64(input.FactWeight),
		NewMinutes: int64(input.FactMinutes),
		NewMeters:  int64(input.FactMeters),
	})
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}

func (s *serviceImpl) validateAccessToExercise(w http.ResponseWriter, chatID int64, exerciseID int64) error {
	ex, err := s.container.GetExerciseUC.Execute(exerciseID)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return err
	}

	user, err := s.container.GetUserUC.Execute(chatID)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return err
	}

	// check access
	if ex.Exercise.WorkoutDay.UserID != user.ID {
		http.Error(w, "access denied", http.StatusForbidden)
		return errors.New("no access to set")
	}

	return nil
}
