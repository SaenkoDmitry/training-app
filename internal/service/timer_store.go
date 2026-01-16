package service

import (
	"sync"

	"github.com/google/uuid"
)

type TimerStore struct {
	userTimers map[int64]map[string]struct{}
	mu         *sync.Mutex
}

func NewTimerStore() *TimerStore {
	return &TimerStore{
		userTimers: make(map[int64]map[string]struct{}),
		mu:         &sync.Mutex{},
	}
}

func (t *TimerStore) NewTimer(chatID int64) string {
	t.mu.Lock()
	defer t.mu.Unlock()
	newTimer, _ := uuid.NewUUID()
	if _, ok := t.userTimers[chatID]; !ok {
		t.userTimers[chatID] = make(map[string]struct{})
	}
	t.userTimers[chatID][newTimer.String()] = struct{}{}
	return newTimer.String()
}

func (t *TimerStore) HasTimer(chatID int64, timerID string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	if v, ok := t.userTimers[chatID]; ok {
		if _, ok2 := v[timerID]; ok2 {
			return true
		}
	}
	return false
}

func (t *TimerStore) StopTimer(chatID int64, timerID string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	if v, ok := t.userTimers[chatID]; ok {
		if _, ok2 := v[timerID]; ok2 {
			delete(t.userTimers[chatID], timerID)
			return true
		}
	}
	return false
}
