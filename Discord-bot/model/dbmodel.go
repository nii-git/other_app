package model

import (
	"time"
)

// Transaction テーブルの構造体
type Transaction struct {
	UserID    string     `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// StudyTimes テーブルの構造体
type StudyTimes struct {
	StudyTimeMinutes int        `json:"study_time_minutes"`
	UserID           string     `json:"user_id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
}
