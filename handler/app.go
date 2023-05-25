package handler

import (
	"github.com/Group-8-H8/fp-1/database"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	db := database.GetDbInstance()

	_ = db

	route := gin.Default()

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": "success",
		})
	})

	route.Run(":8080")
}
