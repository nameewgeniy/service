package entity

import "time"

type Idea struct {
	ID int `form:"id"`
	Title string `form:"title"`
	CreatedAt time.Time
}
