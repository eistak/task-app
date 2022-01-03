package models

import "time"

// Task is parent table of book_comment
type Task struct {
	ID        int       `json:"id" param:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
