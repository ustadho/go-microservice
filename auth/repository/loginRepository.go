package repository

import "github.com/ustadho/go-microservice/auth/models"

var user []models.User

type LoginRepository struct {
}

func Init() *LoginRepository {
	user = make([]models.User, 0, 2)
	user = append(user,
		models.User{
			TenantID: 1,
			Id:       1,
			Name:     "Ustadho",
			UserName: "ustadho",
			Password: "12345678",
			Roles:    []int{1, 2, 3},
		},
		models.User{
			TenantID: 2,
			Id:       2,
			Name:     "Ana",
			UserName: "ana",
			Password: "12345678",
			Roles:    []int{4},
		},
	)
	return &LoginRepository{}
}

func (l *LoginRepository) GetUserByUserName(userName, password string) (models.User, *models.ErrorDetail) {
	for _, value := range user {
		if value.UserName == userName && value.Password == password {
			return value, nil
		}
	}

	return models.User{}, &models.ErrorDetail{
		ErrorType:    models.ErrorTypeError,
		ErrorMessage: "UserName and password not valid",
	}
}
