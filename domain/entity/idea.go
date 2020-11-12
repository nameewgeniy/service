package entity

import "time"

type Idea struct {
	ID int
	Title string `form:"title" binding:"required"`
	CreatedAt time.Time
}
