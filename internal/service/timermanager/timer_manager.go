package timermanager

import (
	"errors"
	"fmt"
	"github.com/SaenkoDmitry/training-tg-bot/internal/service/push"
	"gorm.io/gorm"
	"sync"
	"time"

	"github.com/SaenkoDmitry/training-tg-bot/internal/models"
)

type TimerManager struct {
	db     *gorm.DB
	mu     sync.Mutex
	timers map[int64]*time.Timer
	push   *push.Service
}

func NewTimerManager(db *gorm.DB, push *push.Service) *TimerManager {
	return &TimerManager{
		db:     db,
		push:   push,
		timers: make(map[int64]*time.Timer),
	}
}

func (tm *TimerManager) Start(userID, workoutID int64, seconds int) (*TimerDTO, error) {
	timer := models.RestTimer{
		UserID:    userID,
		WorkoutID: workoutID,
		EndsAt:    time.Now().Add(time.Duration(seconds) * time.Second),
	}

	if err := tm.db.Create(&timer).Error; err != nil {
		return nil, err
	}

	tm.schedule(&timer)

	return mapTimerDTO(timer), nil
}

func mapTimerDTO(timer models.RestTimer) *TimerDTO {
	return &TimerDTO{
		ID:        timer.ID,
		UserID:    timer.UserID,
		WorkoutID: timer.WorkoutID,
		EndsAt:    timer.EndsAt,
		Canceled:  timer.Canceled,
		Sent:      timer.Sent,
		CreatedAt: timer.CreatedAt,
	}
}

type TimerDTO struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	WorkoutID int64     `json:"workout_id"`
	EndsAt    time.Time `json:"ends_at"`
	Canceled  bool      `json:"canceled"`
	Sent      bool      `json:"sent"`
	CreatedAt time.Time `json:"created_at"`
}

func (tm *TimerManager) Restore() error {
	var timers []models.RestTimer

	if err := tm.db.
		Where("canceled = ? AND sent = ?", false, false).
		Find(&timers).Error; err != nil {
		return err
	}

	for _, t := range timers {
		tm.schedule(&t)
	}

	return nil
}

func (tm *TimerManager) schedule(t *models.RestTimer) {
	duration := time.Until(t.EndsAt)
	if duration < 0 {
		duration = 0
	}

	timer := time.AfterFunc(duration, func() {
		tm.fire(t.ID)
	})

	tm.mu.Lock()
	tm.timers[t.ID] = timer
	tm.mu.Unlock()
}

func (tm *TimerManager) fire(timerID int64) {
	var timer models.RestTimer

	if err := tm.db.First(&timer, timerID).Error; err != nil {
		return
	}

	if timer.Canceled || timer.Sent {
		return
	}

	// отправляем push
	if err := tm.push.SendWorkoutFinished(timer.UserID, timer.WorkoutID); err != nil {
		fmt.Println("push error:", err)
		return
	}

	timer.Sent = true
	tm.db.Save(&timer)

	tm.mu.Lock()
	delete(tm.timers, timerID)
	tm.mu.Unlock()
}

func (tm *TimerManager) Cancel(timerID int64, userID int64) error {
	var timer models.RestTimer

	if err := tm.db.First(&timer, timerID).Error; err != nil {
		return err
	}

	if timer.UserID != userID {
		return errors.New("forbidden")
	}

	timer.Canceled = true
	tm.db.Save(&timer)

	tm.mu.Lock()
	if t, ok := tm.timers[timerID]; ok {
		t.Stop()
		delete(tm.timers, timerID)
	}
	tm.mu.Unlock()

	return nil
}
