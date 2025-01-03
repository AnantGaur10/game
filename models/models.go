package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"unique;size:15;not null;min=6,max=15"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"size:70;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}
