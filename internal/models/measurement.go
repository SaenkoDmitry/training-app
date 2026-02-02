package models

import "time"

type Measurement struct {
	ID        int64 `gorm:"primaryKey;autoIncrement"`
	UserID    int64
	CreatedAt time.Time

	Shoulders int
	Chest     int
	HandLeft  int
	HandRight int
	Waist     int
	Buttocks  int
	HipLeft   int
	HipRight  int
	CalfLeft  int
	CalfRight int
	Weight    int
}

func (*Measurement) TableName() string {
	return "measurements"
}
