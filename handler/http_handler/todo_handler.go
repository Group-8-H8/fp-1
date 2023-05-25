package http_handler

import (
	"fmt"
	"net/http"

	"github.com/Group-8-H8/fp-1/dto"
	"github.com/Group-8-H8/fp-1/pkg/errs"
	"github.com/Group-8-H8/fp-1/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TodoHandler interface {
	CreateTodo(ctx *gin.Context)
}

type todoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) TodoHandler {
	return &todoHandler{todoService: todoService}
}

func (t *todoHandler) CreateTodo(ctx *gin.Context) {
	var requestBody dto.CreateTodoRequest

	if errBindJSON := ctx.ShouldBindJSON(&requestBody); errBindJSON != nil {
		errCasting, ok := errBindJSON.(validator.ValidationErrors)
		if !ok {
			errBind := errs.NewBadRequest("invalid request body")
			ctx.AbortWithStatusJSON(errBind.Status(), errBind)
			return
		}
		errBinds := []string{}
		for _, e := range errCasting {
			errBind := fmt.Sprintf("error on field %s, condition : %s", e.Field(), e.ActualTag())
			errBinds = append(errBinds, errBind)
		}
		NewErrBinds := errs.NewUnprocessableEntity(errBinds)
		ctx.AbortWithStatusJSON(NewErrBinds.Status(), NewErrBinds)
		return
	}

	response, errCreated := t.todoService.CreateTodo(requestBody)
	if errCreated != nil {
		ctx.AbortWithStatusJSON(errCreated.Status(), errCreated)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
