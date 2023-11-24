package models

import "time"

type Article struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(255);NOT NULL" json:"title" binding:"required"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Articles []Article
