package main

import (
	"github.com/aZ4ziL/web-kafekoding/controllers"
	"github.com/aZ4ziL/web-kafekoding/models"
	"github.com/gin-contrib/sessions"
	gormsession "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.HTMLRender = createMyRender()

	router.Static("/static", "./static")

	store := gormsession.NewStore(models.GetDB(), true, []byte("MySecretKey"))
	router.Use(sessions.Sessions("kafekodingID", store))

	// user
	userGroup := router.Group("/user")
	controllers.UserController(userGroup)

	router.Run(":8000")
}
