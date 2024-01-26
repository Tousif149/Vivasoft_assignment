package domain

import (
	"book-crud/pkg/models"
	"book-crud/pkg/types"
)

type IUserRepo interface {
	CreateUser(user *models.UserDetail) error
	GetUser(username *string) (*models.UserDetail, error)
}
type IAuthService interface {
	LoginUser(user *types.LoginRequest)(string,error)
	SignupUser(user *types.UserRequest) error
}
