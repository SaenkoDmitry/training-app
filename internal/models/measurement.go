package models

import "time"

type Measurement struct {
	ID        int64 `gorm:"primaryKey;autoIncrement"`
	UserID    int64
	CreatedAt time.Time

	shoulders int
	chest     int
	hands     int
	waist     int
	buttocks  int
	hip       int
	Calves    int
	weight    int
}

func (*Measurement) TableName() string {
	return "measurements"
}
