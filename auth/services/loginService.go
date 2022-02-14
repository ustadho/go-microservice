package services

import (
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ustadho/go-microservice/auth/models"
	"github.com/ustadho/go-microservice/auth/repository"
	"github.com/ustadho/go-microservice/auth/token"
)

type LoginService struct {
	logger          *log.Logger
	flags           *models.Flags
	loginRepository *repository.LoginRepository
}

func NewLoginService(l *log.Logger, f *models.Flags) *LoginService {
	return &LoginService{
		logger:          l,
		flags:           f,
		loginRepository: repository.Init(),
	}
}

func (l *LoginService) GetToken(loginModel models.LoginRequest, origin string) (string, *models.ErrorDetail) {
	user, err := l.loginRepository.GetUserByUserName(loginModel.Username, loginModel.Password)
	if err != nil {
		return "", nil
	}

	var claims = &models.JwtClaims{
		TenantID: strconv.Itoa(user.TenantID),
		Username: user.UserName,
		Roles:    user.Roles,
		StandardClaims: jwt.StandardClaims{
			Audience: origin,
		},
	}
	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(2) * time.Hour)
	return token.GenerateToken(claims, expirationTime)
}

func (s *LoginService) VerifyToken(tokenString, referer string) (bool, *models.JwtClaims) {
	return token.VerifyToken(tokenString, referer)
}
