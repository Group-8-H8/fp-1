package todo_pg

import (
	"github.com/Group-8-H8/fp-1/entity"
	"github.com/Group-8-H8/fp-1/pkg/errs"
	"github.com/Group-8-H8/fp-1/repository/todo_repository"
	"gorm.io/gorm"
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
