package dto

import "time"

type TodoRequest struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

type TodoResponse struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoCreateResponse struct {
	TodoResponse
	CreatedAt time.Time `json:"created_at"`
}

type TodoUpdateResponse struct {
	TodoResponse
	UpdatedAt time.Time `json:"updated_at"`
}

type TodoGetResponse struct {
	TodoResponse
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
