package todo_repository

import (
	"github.com/Group-8-H8/fp-1/entity"
	"github.com/Group-8-H8/fp-1/pkg/errs"
)

type TodoRepository interface {
	CreateTodo(payload entity.Todo) (*entity.Todo, errs.MessageErr)
}
