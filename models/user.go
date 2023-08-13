package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	createdAt time.Time
	updatedAt time.Time
}
