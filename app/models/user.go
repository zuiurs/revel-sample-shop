package models

import (
	"time"
)

type User struct {
	ID             uint64 `json:"id" gorm:"primary_key"`
	Username       string `json:"username" gorm:"size:64"`
	Password       string `json:"-"`
	HashedPassword []byte `json:"password"`
	RoleID         uint64 `json:"role_id"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
