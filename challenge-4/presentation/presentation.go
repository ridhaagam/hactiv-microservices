package presentation

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"-"`
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
}
