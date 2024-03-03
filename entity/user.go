package entity

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	Username    string    `json:"username" gorm:"not null; unique"`
	Password    string    `json:"password" gorm:"not null"`
	HasVoted    bool      `json:"-"`
	IsCandidate bool      `json:"-"`
	Votes       uint      `json:"-"`
}
