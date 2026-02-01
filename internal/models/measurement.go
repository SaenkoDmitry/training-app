package models

import "time"

type Measurement struct {
	ID        int64 `gorm:"primaryKey;autoIncrement"`
	UserID    int64
	CreatedAt time.Time

	Shoulders int
	Chest     int
	Hands     int
	Waist     int
	Buttocks  int
	Hips      int
	Calves    int
	Weight    int
}

func (*Measurement) TableName() string {
	return "measurements"
}
