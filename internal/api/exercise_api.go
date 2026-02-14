package api

import (
	"github.com/SaenkoDmitry/training-tg-bot/internal/api/helpers"
	"github.com/SaenkoDmitry/training-tg-bot/internal/middlewares"
	"net/http"
)

func (s *serviceImpl) DeleteExercise(w http.ResponseWriter, r *http.Request) {
	claims, ok := middlewares.FromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	exerciseID, err := helpers.ParseInt64Param("id", w, r)
	if err != nil {
		return
	}

	err = s.validateAccessToExercise(w, claims.ChatID, exerciseID)
	if err != nil {
		return
	}

	_, err = s.container.DeleteExerciseUC.Execute(exerciseID)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}
