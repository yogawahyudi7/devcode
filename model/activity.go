package model

import "time"

type Activity struct {
	ActivityId uint `gorm:"primarykey"`
	Title      string
	Email      *string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
}

func (Activity) TableName() string {
	return "activities"
}
