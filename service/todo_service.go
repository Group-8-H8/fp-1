package service

import (
	"time"

	"github.com/Group-8-H8/fp-1/dto"
	"github.com/Group-8-H8/fp-1/entity"
	"github.com/Group-8-H8/fp-1/pkg/errs"
	"github.com/Group-8-H8/fp-1/repository/todo_repository"
)

type TodoService interface {
	CreateTodo(payload dto.CreateTodoRequest) (*dto.Response, errs.MessageErr)
}

type todoService struct {
	todoRepo todo_repository.TodoRepository
}

func NewTodoService(todoRepo todo_repository.TodoRepository) TodoService {
	return &todoService{todoRepo: todoRepo}
}

func (t *todoService) CreateTodo(payload dto.CreateTodoRequest) (*dto.Response, errs.MessageErr) {
	todo := entity.Todo{
		Title:     payload.Title,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdTodo, err := t.todoRepo.CreateTodo(todo)
	if err != nil {
		return nil, err
	}

	todoResponse := dto.TodoResponse{
		Id:        createdTodo.Id,
		Title:     createdTodo.Title,
		Completed: createdTodo.Completed,
	}

	todoGetResponse := dto.TodoCreateResponse{
		TodoResponse: todoResponse,
		CreatedAt:    createdTodo.CreatedAt,
	}

	response := &dto.Response{
		Status:  "CREATED",
		Message: "todo created successfully",
		Data:    todoGetResponse,
	}

	return response, nil
}
