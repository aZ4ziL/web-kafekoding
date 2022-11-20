package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aZ4ziL/web-kafekoding/auth"
	"github.com/aZ4ziL/web-kafekoding/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type UserHandler interface {
	Login() gin.HandlerFunc
	Register() gin.HandlerFunc
}

type userHandler struct{}

func NewUserHandler() UserHandler {
	return &userHandler{}
}

func (u userHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		if user := session.Get("user"); user != nil {
			ctx.Redirect(http.StatusFound, "/")
			return
		}

		if ctx.Request.Method == "GET" {
			defer flasher.Del()

			ctx.HTML(http.StatusOK, "login", gin.H{
				"flasher": flasher,
			})
		}

		if ctx.Request.Method == "POST" {
			studentIdNumber := ctx.PostForm("student_id_number")
			studentIdNumberInt, _ := strconv.Atoi(studentIdNumber)
			password := ctx.PostForm("password")

			user, err := models.NewUserModel().GetUserByStudentIDNumber(uint(studentIdNumberInt))
			if err != nil {
				flasher.Set("danger", "Username or password is incorrect!")
				ctx.Redirect(http.StatusFound, "/user/login")
				return
			}

			if !auth.DecryptionPassword(user.Password, password) {
				flasher.Set("danger", "Username or password is incorrect!")
				ctx.Redirect(http.StatusFound, "/user/login")
				return
			}

			userSession := map[string]interface{}{
				"id":                user.ID,
				"full_name":         user.FullName,
				"student_id_number": user.StudentIDNumber,
				"email":             user.Email,
				"last_login":        user.LastLogin,
				"date_joined":       user.DateJoined,
			}

			session.Set("user", userSession)
			if err := session.Save(); err != nil {
				http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
				return
			}

			ctx.Redirect(http.StatusFound, "/")
			return
		}
	}
}

func (u userHandler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		if user := session.Get("user"); user != nil {
			ctx.Redirect(http.StatusFound, "/")
			return
		}

		if ctx.Request.Method == "GET" {
			defer flasher.Del()
			ctx.HTML(http.StatusOK, "register", gin.H{
				"flasher": flasher,
			})
		}

		if ctx.Request.Method == "POST" {
			var user models.User

			err := ctx.ShouldBindWith(&user, binding.FormPost)
			if err != nil {
				http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
				return
			}

			validate = validator.New()
			err = validate.Struct(&user)
			if err != nil {
				if _, ok := err.(*validator.InvalidValidationError); ok {
					log.Printf("error :%s", err)
					return
				}
				errorMessages := []string{}
				for _, err := range err.(validator.ValidationErrors) {
					errorMessage := fmt.Sprintf("error on field: %s, with error: %s.", err.Field(), err.ActualTag())
					errorMessages = append(errorMessages, errorMessage)
				}
				flasher.Set("danger", errorMessages[0])
				ctx.Redirect(http.StatusFound, "/user/register")
				return
			}

			err = models.NewUserModel().CreateNewUser(&user)
			if err != nil {
				flasher.Set("danger", err.Error())
				ctx.Redirect(http.StatusFound, "/user/register")
				return
			}

			flasher.Set("success", fmt.Sprintf("Berhasil mendaftarkan akun dengan nomor induk mahasiswa: %d.", user.StudentIDNumber))
			ctx.Redirect(http.StatusFound, "/user/login")
			return
		}
	}
}
