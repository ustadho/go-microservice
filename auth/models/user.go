package models

type User struct {
	TenantID int
	Id       int
	Name     string
	UserName string
	Password string
	Roles    []int
}
