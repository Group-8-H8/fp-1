package service

import (
	"time"

	"github.com/Group-8-H8/fp-1/dto"
	"github.com/Group-8-H8/fp-1/entity"
	"github.com/Group-8-H8/fp-1/pkg/errs"
	"github.com/Group-8-H8/fp-1/repository/todo_repository"
)

type TodoService interface {
	CreateTodo(payload dto.TodoRequest) (*dto.Response, errs.MessageErr)
	GetTodos() (*dto.Response, errs.MessageErr)
	GetTodo(todoId int) (*dto.Response, errs.MessageErr)
	UpdateTodo(todoId int, payload dto.TodoRequest) (*dto.Response, errs.MessageErr)
	DeleteTodo(todoId int) (*dto.Response, errs.MessageErr)
}

type todoService struct {
	todoRepo todo_repository.TodoRepository
}

func NewTodoService(todoRepo todo_repository.TodoRepository) TodoService {
	return &todoService{todoRepo: todoRepo}
}

func (t *todoService) CreateTodo(payload dto.TodoRequest) (*dto.Response, errs.MessageErr) {
	todo := entity.Todo{
		Title:     payload.Title,
		Completed: payload.Completed,
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

	todoCreateResponse := dto.TodoCreateResponse{
		TodoResponse: todoResponse,
		CreatedAt:    createdTodo.CreatedAt,
	}

	response := &dto.Response{
		Status:  "CREATED",
		Message: "todo created successfully",
		Data:    todoCreateResponse,
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

func (t *todoService) UpdateTodo(todoId int, payload dto.TodoRequest) (*dto.Response, errs.MessageErr) {
	if _, errCheck := t.todoRepo.GetTodoById(todoId); errCheck != nil && errCheck.Code() == 404 {
		return nil, errCheck
	}

	todo := entity.Todo{
		Id:        uint(todoId),
		Title:     payload.Title,
		Completed: payload.Completed,
		UpdatedAt: time.Now(),
	}

	updatedTodo, err := t.todoRepo.UpdateTodo(todo)
	if err != nil {
		return nil, err
	}

	todoResponse := dto.TodoResponse{
		Id:        updatedTodo.Id,
		Title:     updatedTodo.Title,
		Completed: updatedTodo.Completed,
	}

	todoUpdateResponse := dto.TodoUpdateResponse{
		TodoResponse: todoResponse,
		UpdatedAt:    updatedTodo.UpdatedAt,
	}

	response := &dto.Response{
		Status:  "OK",
		Message: "todo updated successfully",
		Data:    todoUpdateResponse,
	}

	return response, nil
}

func (t *todoService) DeleteTodo(todoId int) (*dto.Response, errs.MessageErr) {
	if _, errCheck := t.todoRepo.GetTodoById(todoId); errCheck != nil && errCheck.Code() == 404 {
		return nil, errCheck
	}

	if errDeleted := t.todoRepo.DeleteTodo(todoId); errDeleted != nil {
		return nil, errDeleted
	}

	response := &dto.Response{
		Status:  "OK",
		Message: "todo deleted successfully",
	}

	return response, nil
}
