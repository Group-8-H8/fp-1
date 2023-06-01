package dto

import "time"

type TodoRequest struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

type TodoResponse struct {
	Id        uint       `json:"id"`
	Title     string     `json:"title"`
	Completed bool       `json:"completed"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrData struct {
	ErrStatus  string `json:"status"`
	ErrMessage any    `json:"message"`
	ErrCode    int    `json:"code"`
}
