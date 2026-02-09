package helpers

import (
	"net/http"
	"strconv"
)

func GetOffsetLimit(r *http.Request, defaultLimit, maxLimit int) (int, int) {
	q := r.URL.Query()

	offset, _ := strconv.Atoi(q.Get("offset"))
	limit, _ := strconv.Atoi(q.Get("limit"))

	if offset < 0 {
		offset = 0
	}

	if limit <= 0 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}
	return offset, limit
}

func ParseInt64Param(name string, w http.ResponseWriter, r *http.Request) (int64, error) {
	entityIDStr := r.PathValue(name)
	entityID, err := strconv.ParseInt(entityIDStr, 10, 64)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return 0, err
	}
	return entityID, nil
}
