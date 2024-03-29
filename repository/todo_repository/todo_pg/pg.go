package todo_pg

import (
	"errors"

	"github.com/Group-8-H8/fp-1/entity"
	"github.com/Group-8-H8/fp-1/pkg/errs"
	"github.com/Group-8-H8/fp-1/repository/todo_repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type todoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) todo_repository.TodoRepository {
	return &todoRepo{db: db}
}

func (t *todoRepo) CreateTodo(payload entity.Todo) (*entity.Todo, errs.MessageErr) {
	if err := t.db.Create(&payload).Error; err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &payload, nil
}

func (t *todoRepo) GetAllTodos() ([]entity.Todo, errs.MessageErr) {
	var todos []entity.Todo

	if err := t.db.Find(&todos).Error; err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return todos, nil
}

func (t *todoRepo) GetTodoById(todoId int) (*entity.Todo, errs.MessageErr) {
	var todo entity.Todo

	if err := t.db.First(&todo, todoId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("todo not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &todo, nil
}

func (t *todoRepo) UpdateTodo(payload entity.Todo) (*entity.Todo, errs.MessageErr) {
	err := t.db.Model(&payload).Where("id = ?", payload.Id).Clauses(clause.Returning{}).Updates(entity.Todo{Title: payload.Title, Completed: payload.Completed, UpdatedAt: payload.UpdatedAt}).Error
	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &payload, nil
}

func (t *todoRepo) DeleteTodo(todoId int) errs.MessageErr {
	if err := t.db.Delete(&entity.Todo{}, todoId).Error; err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
