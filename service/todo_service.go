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
	GetTodos() (*dto.Response, errs.MessageErr)
	GetTodo(todoId int) (*dto.Response, errs.MessageErr)
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

func (t *todoService) GetTodos() (*dto.Response, errs.MessageErr) {
	getTodos, err := t.todoRepo.GetAllTodos()
	if err != nil {
		return nil, err
	}

	todoGetResponses := []dto.TodoGetResponse{}
	for _, todo := range getTodos {
		todoResponse := dto.TodoResponse{
			Id:        todo.Id,
			Title:     todo.Title,
			Completed: todo.Completed,
		}

		todoGetResponse := dto.TodoGetResponse{
			TodoResponse: todoResponse,
			CreatedAt:    todo.CreatedAt,
			UpdatedAt:    todo.UpdatedAt,
		}

		todoGetResponses = append(todoGetResponses, todoGetResponse)
	}

	response := &dto.Response{
		Status:  "OK",
		Message: "all todos found",
		Data:    todoGetResponses,
	}

	return response, nil
}

func (t *todoService) GetTodo(todoId int) (*dto.Response, errs.MessageErr) {
	getTodo, err := t.todoRepo.GetTodoById(todoId)
	if err != nil {
		return nil, err
	}

	todoResponse := dto.TodoResponse{
		Id:        getTodo.Id,
		Title:     getTodo.Title,
		Completed: getTodo.Completed,
	}

	todoGetResponse := dto.TodoGetResponse{
		TodoResponse: todoResponse,
		CreatedAt:    getTodo.CreatedAt,
		UpdatedAt:    getTodo.UpdatedAt,
	}

	response := &dto.Response{
		Status:  "OK",
		Message: "todo found",
		Data:    todoGetResponse,
	}

	return response, nil
}
