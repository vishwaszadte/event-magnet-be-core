package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name"`
	Photo    string    `json:"photo"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"password" gorm:"not null"`
	Bio      string    `json:"bio"`
}
