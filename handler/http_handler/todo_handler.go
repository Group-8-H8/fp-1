package http_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Group-8-H8/fp-1/dto"
	"github.com/Group-8-H8/fp-1/pkg/errs"
	"github.com/Group-8-H8/fp-1/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TodoHandler interface {
	CreateTodo(ctx *gin.Context)
	GetTodos(ctx *gin.Context)
	GetTodo(ctx *gin.Context)
	UpdateTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
}

type todoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) TodoHandler {
	return &todoHandler{todoService: todoService}
}

func (t *todoHandler) CreateTodo(ctx *gin.Context) {
	var requestBody dto.TodoRequest

	if errBindJSON := ctx.ShouldBindJSON(&requestBody); errBindJSON != nil {
		errCasting, ok := errBindJSON.(validator.ValidationErrors)
		if !ok {
			errBind := errs.NewBadRequest("invalid request body")
			ctx.AbortWithStatusJSON(errBind.Code(), errBind)
			return
		}
		errBinds := []string{}
		for _, e := range errCasting {
			errBind := fmt.Sprintf("error on field %s, condition : %s", e.Field(), e.ActualTag())
			errBinds = append(errBinds, errBind)
		}
		NewErrBinds := errs.NewUnprocessableEntity(errBinds)
		ctx.AbortWithStatusJSON(NewErrBinds.Code(), NewErrBinds)
		return
	}

	response, errCreated := t.todoService.CreateTodo(requestBody)
	if errCreated != nil {
		ctx.AbortWithStatusJSON(errCreated.Code(), errCreated)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (t *todoHandler) GetTodos(ctx *gin.Context) {
	response, errGetTodos := t.todoService.GetTodos()
	if errGetTodos != nil {
		ctx.AbortWithStatusJSON(errGetTodos.Code(), errGetTodos)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (t *todoHandler) GetTodo(ctx *gin.Context) {
	param := ctx.Param("todoId")
	todoId, errConv := strconv.Atoi(param)
	if errConv != nil {
		errConv := errs.NewBadRequest("invalid todo's id")
		ctx.AbortWithStatusJSON(errConv.Code(), errConv)
		return
	}

	response, errGetTodos := t.todoService.GetTodo(todoId)
	if errGetTodos != nil {
		ctx.AbortWithStatusJSON(errGetTodos.Code(), errGetTodos)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (t *todoHandler) UpdateTodo(ctx *gin.Context) {
	param := ctx.Param("todoId")
	todoId, errConv := strconv.Atoi(param)
	if errConv != nil {
		newErrConv := errs.NewBadRequest("invalid todo's id")
		ctx.AbortWithStatusJSON(newErrConv.Code(), newErrConv)
		return
	}

	var requestBody dto.TodoRequest

	if errBinding := ctx.ShouldBindJSON(&requestBody); errBinding != nil {
		errCasting, ok := errBinding.(validator.ValidationErrors)
		if !ok {
			newErrBind := errs.NewBadRequest("invalid request body")
			ctx.AbortWithStatusJSON(newErrBind.Code(), newErrBind)
			return
		}

		errBinds := []string{}
		for _, e := range errCasting {
			errBind := fmt.Sprintf("error on field %s, condition : %s", e.Field(), e.ActualTag())
			errBinds = append(errBinds, errBind)
		}
		newErrBind := errs.NewUnprocessableEntity(errBinds)
		ctx.AbortWithStatusJSON(newErrBind.Code(), newErrBind)
		return
	}

	response, errUpdated := t.todoService.UpdateTodo(todoId, requestBody)
	if errUpdated != nil {
		ctx.AbortWithStatusJSON(errUpdated.Code(), errUpdated)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (t *todoHandler) DeleteTodo(ctx *gin.Context) {
	param := ctx.Param("todoId")
	todoId, errConv := strconv.Atoi(param)
	if errConv != nil {
		newErrConv := errs.NewBadRequest("invalid todo's id")
		ctx.AbortWithStatusJSON(newErrConv.Code(), newErrConv)
		return
	}

	response, errDeleted := t.todoService.DeleteTodo(todoId)
	if errDeleted != nil {
		ctx.AbortWithStatusJSON(errDeleted.Code(), errDeleted)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
