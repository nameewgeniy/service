package dto

type ListQuery struct {
	Page int `form:"page"`
	Count int `form:"count"`
}
