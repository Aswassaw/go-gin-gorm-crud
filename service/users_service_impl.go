package service

import (
	"go-gin-gorm-crud/data/request"
	"go-gin-gorm-crud/data/response"
	"go-gin-gorm-crud/helper"
	"go-gin-gorm-crud/model"
	"go-gin-gorm-crud/repository"

	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewUsersServiceImpl(usersRepository repository.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}

func (u *UsersServiceImpl) Create(user request.CreateUserRequest) {
	err := u.Validate.Struct(user)
	helper.PanicIfError(err)

	userModel := model.User{
		Name:      user.Name,
		Age:       user.Age,
	}
	u.UsersRepository.Create(userModel)
}

func (u *UsersServiceImpl) FindAll() []response.UserResponse {
	result := u.UsersRepository.FindAll()

	var users []response.UserResponse
	for _, value := range result {
		user := response.UserResponse{
			Id:        value.Id,
			Name:      value.Name,
			Age:       value.Age,
		}
		users = append(users, user)
	}

	return users
}

func (u *UsersServiceImpl) FindById(userId int) response.UserResponse {
	userData, err := u.UsersRepository.FindById(userId)
	helper.PanicIfError(err)

	userResponse := response.UserResponse{
		Id:        userData.Id,
		Name:      userData.Name,
		Age:       userData.Age,
	}
	return userResponse
}

func (u *UsersServiceImpl) Update(user request.UpdateUserRequest) {
	userData, err := u.UsersRepository.FindById(user.Id)
	helper.PanicIfError(err)

	userData.Name = user.Name
	userData.Age = user.Age

	u.UsersRepository.Update(userData)
}

func (u *UsersServiceImpl) Delete(userId int) {
	u.UsersRepository.Delete(userId)
}
