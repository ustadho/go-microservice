package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ustadho/go-microservice/auth/models"
	"github.com/ustadho/go-microservice/auth/services"
)

type LoginHandler struct {
	logger       *log.Logger
	flags        *models.Flags
	loginService *services.LoginService
}

func NewLogin(l *log.Logger, f *models.Flags) *LoginHandler {
	loginService := services.NewLoginService(l, f)
	return &LoginHandler{
		logger:       l,
		flags:        f,
		loginService: loginService,
	}
}
func (l *LoginHandler) Login(ctx *gin.Context) {
	var loginObj models.LoginRequest
	if err := ctx.ShouldBindJSON(&loginObj); err != nil {
		var errors []models.ErrorDetail = make([]models.ErrorDetail, 0, 1)
		errors = append(errors, models.ErrorDetail{
			ErrorType:    models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		})
		badRequest(ctx, http.StatusBadRequest, "Invalid request", errors)
	}
	tokenString, err := l.loginService.GetToken(loginObj, ctx.Request.Header.Get("Referer"))

	if err != nil {
		badRequest(ctx, http.StatusBadRequest, "error generating token", []models.ErrorDetail{
			*err,
		})
	}
	ok(ctx, http.StatusOK, "token created", tokenString)
}

func (l *LoginHandler) VerifyToken(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("apiKey")
	referer := ctx.Request.Header.Get("Referer")

	valid, claims := l.loginService.VerifyToken(tokenString, referer)
	if !valid {
		returnUnauthorized(ctx)
		return
	}
	ok(ctx, http.StatusOK, "token is valid", claims)
}
