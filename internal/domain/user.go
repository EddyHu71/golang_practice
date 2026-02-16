package domain

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
