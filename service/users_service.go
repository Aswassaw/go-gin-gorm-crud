package service

import (
	"go-gin-gorm-crud/data/request"
	"go-gin-gorm-crud/data/response"
)

type UsersService interface {
	Create(user request.CreateUserRequest)
	FindAll() []response.UserResponse
	FindById(userId int) response.UserResponse
	Update(user request.UpdateUserRequest)
	Delete(userId int)
}
