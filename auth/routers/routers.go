package routers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ustadho/go-microservice/auth/handlers"
	"github.com/ustadho/go-microservice/auth/models"
	"github.com/ustadho/go-microservice/auth/token"
)

type AuthRouter struct {
	logger       *log.Logger
	loginHandler *handlers.LoginHandler
	flags        *models.Flags
}

func NewRoute(l *log.Logger, f *models.Flags) *AuthRouter {
	loginHandler := handlers.NewLogin(l, f)
	token.Init()

	return &AuthRouter{
		logger:       l,
		flags:        f,
		loginHandler: loginHandler,
	}
}

func (r *AuthRouter) RegisterRoutes() *gin.Engine {
	ginEngine := gin.Default()
	ginEngine.POST("/login", r.loginHandler.Login)
	ginEngine.POST("/verifyToken", r.loginHandler.VerifyToken)
	return ginEngine
}
