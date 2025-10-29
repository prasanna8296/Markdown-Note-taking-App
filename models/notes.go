package models

import "time"

type Note struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `json:"title"`
	Filename  string    `json:"filename"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
