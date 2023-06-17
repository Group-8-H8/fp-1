package handler

import (
	"github.com/Group-8-H8/fp-1/database"
	_ "github.com/Group-8-H8/fp-1/docs"
	"github.com/Group-8-H8/fp-1/handler/http_handler"
	"github.com/Group-8-H8/fp-1/repository/todo_repository/todo_pg"
	"github.com/Group-8-H8/fp-1/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Final Project 1 - Group 8 Hacktiv8
// @version         1.0
// @description     This is a documentation for todolist API from final project 1 - Group 8 Hacktiv8

// @host      fp-1-production.up.railway.app
// @BasePath  /todos

func StartApp() {
	db := database.GetDbInstance()

	todoRepo := todo_pg.NewTodoRepo(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := http_handler.NewTodoHandler(todoService)

	route := gin.Default()

	todoRoute := route.Group("/todos")
	{
		todoRoute.POST("/", todoHandler.CreateTodo)
		todoRoute.GET("/", todoHandler.GetTodos)
		todoRoute.GET("/:todoId", todoHandler.GetTodo)
		todoRoute.PUT("/:todoId", todoHandler.UpdateTodo)
		todoRoute.DELETE("/:todoId", todoHandler.DeleteTodo)
	}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route.Run(":8080")
}
