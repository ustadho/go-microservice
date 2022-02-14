package models

type LoginRequest struct {
	Username   string `json:"username" form:"username" validate:"required"`
	Password   string `json:"password" form:"password" validate:"required"`
	RememberMe bool   `json:"rememberMe" form:"rememberMe"`
}
