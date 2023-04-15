package model

import (
	"time"
)

type Todo struct {
	// gorm.Model
	TodoId          uint `gorm:"primarykey"`
	ActivityGroupId int
	Title           string
	Priority        string
	IsActive        bool
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	DeletedAt       *time.Time
}
