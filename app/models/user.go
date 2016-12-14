package models

import (
	"time"
)

type User struct {
	ID        uint64     `json:"id" gorm:"primary_key"`
	Name      string     `json:"name" gorm:"size:64" validate:"max=64"`
	Role      Role       `json:"role" gorm:"ForeignKey:ID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
