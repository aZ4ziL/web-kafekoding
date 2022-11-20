package controllers

import (
	"github.com/aZ4ziL/web-kafekoding/handlers"
	"github.com/gin-gonic/gin"
)

func UserController(group *gin.RouterGroup) {
	userHandler := handlers.NewUserHandler()

	group.GET("/login", userHandler.Login())
	group.POST("/login", userHandler.Login())
	group.GET("/register", userHandler.Register())
	group.POST("/register", userHandler.Register())
}
