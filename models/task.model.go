package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Title       string    `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Description string    `gorm:"not null" json:"description,omitempty"`
	IsDone      bool      `gorm:"not null" json:"isdone,omitempty"`
	User        uuid.UUID `gorm:"not null" json:"user,omitempty"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateTaskRequest struct {
	Title       string    `json:"title"  binding:"required"`
	Description string    `json:"description" binding:"required"`
	IsDone      bool      `json:"isdone,omitempty"`
	User        string    `json:"user,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type UpdateTask struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	IsDone      bool      `json:"isdone,omitempty"`
	User        string    `json:"user,omitempty"`
	CreateAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type DeleteTask struct {
	ID uuid.UUID `json:"id,omitempty"`
}
