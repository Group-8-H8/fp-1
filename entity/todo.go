package entity

import "time"

type Todo struct {
	Id        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
