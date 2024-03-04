package model

import (
	"time"

	"github.com/google/uuid"
)

type PostCreate struct {
	ID        uuid.UUID `json:"-"`
	UserID    uuid.UUID `json:"-"`
	Title     string    `json:"title" binding:"required,min=1"`
	Body      string    `json:"body" binding:"required,min=1"`
	CreatedAt time.Time `json:"-"`
}

type PostParam struct {
	ID     uuid.UUID `json:"-"`
	UserID uuid.UUID `json:"-"`
}
