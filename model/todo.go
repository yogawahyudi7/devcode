package model

import (
	"time"
)

type Todo struct {
	TodoId          uint `gorm:"primarykey"`
	ActivityGroupId int
	Title           *string
	Priority        *string
	IsActive        *bool
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	DeletedAt       *time.Time
}

func (Todo) TableName() string {
	return "todos"
}
