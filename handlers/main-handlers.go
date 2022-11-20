package handlers

import (
	"net/http"

	"github.com/aZ4ziL/web-kafekoding/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type MainHandler interface {
	Index() gin.HandlerFunc
}

type mainHandler struct{}

func NewMainHandler() MainHandler {
	return &mainHandler{}
}

func (m mainHandler) Index() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		user := session.Get("user")

		classes := models.NewClassModel().GetAllClass()

		ctx.HTML(http.StatusOK, "index", gin.H{
			"user":    user,
			"classes": classes,
		})
	}
}
