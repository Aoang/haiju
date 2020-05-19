package model

import (
	"time"
)

// Model represents meta data of entity.
type Model struct {
	ID        uint64     `gorm:"primary_key" json:"-"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
