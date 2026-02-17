package api

import (
	"encoding/json"
	"errors"
	"github.com/SaenkoDmitry/training-tg-bot/internal/api/helpers"
	"github.com/SaenkoDmitry/training-tg-bot/internal/middlewares"
	"gorm.io/gorm"
	"net/http"
)

func (s *serviceImpl) StartTimer(w http.ResponseWriter, r *http.Request) {
	claims, ok := middlewares.FromContext(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var req struct {
		WorkoutID int64 `json:"workout_id"`
		Seconds   int   `json:"seconds"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := s.container.GetUserUC.Execute(claims.ChatID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	timer, err := s.timerManager.Start(user.ID, req.WorkoutID, req.Seconds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(timer)
}

func (s *serviceImpl) CancelTimer(w http.ResponseWriter, r *http.Request) {
	claims, ok := middlewares.FromContext(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	timerID, err := helpers.ParseInt64Param("id", w, r)
	if err != nil {
		return
	}

	user, err := s.container.GetUserUC.Execute(claims.ChatID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.timerManager.Cancel(timerID, user.ID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			w.WriteHeader(http.StatusNotFound)
		case err.Error() == "forbidden":
			w.WriteHeader(http.StatusForbidden)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
