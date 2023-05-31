package handler

import (
	"github.com/Group-8-H8/fp-1/database"
	"github.com/Group-8-H8/fp-1/handler/http_handler"
	"github.com/Group-8-H8/fp-1/repository/todo_repository/todo_pg"
	"github.com/Group-8-H8/fp-1/service"
	"github.com/gin-gonic/gin"
)

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
	}

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": "success",
		})
	})

	route.Run(":8080")
}
