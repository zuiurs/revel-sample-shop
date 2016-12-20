package models

import (
	"time"
)

type Merchandise struct {
	ID          uint64     `json:"id" gorm:"primary_key"`
	Name        string     `json:"name" gorm:"size:64"`
	Price       int        `json:"price"`
	Description string     `json:"description" gorm:"size:255"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
