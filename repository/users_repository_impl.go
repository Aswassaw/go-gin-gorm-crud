package repository

import (
	"errors"
	"go-gin-gorm-crud/data/request"
	"go-gin-gorm-crud/helper"
	"go-gin-gorm-crud/model"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	DB *gorm.DB
}

func NewUsersRepositoryImpl(DB *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{DB: DB}
}

func (u *UsersRepositoryImpl) Create(user model.User) {
	result := u.DB.Create(&user)
	helper.PanicIfError(result.Error)
}

func (u *UsersRepositoryImpl) FindAll() []model.User {
	var users []model.User
	result := u.DB.Find(&users)
	helper.PanicIfError(result.Error)

	return users
}

func (u *UsersRepositoryImpl) FindById(userId int) (user model.User, err error) {
	var userById model.User
	result := u.DB.Find(&user, userId)

	if result != nil {
		return user, nil
	}

	return userById, errors.New("user id not found")
}

func (u *UsersRepositoryImpl) Update(user model.User) {
	var updateUser = request.UpdateUserRequest{
		Id:        user.Id,
		Name:      user.Name,
		Age:       user.Age,
	}
	result := u.DB.Model(&user).Updates(updateUser)
	helper.PanicIfError(result.Error)
}

func (u *UsersRepositoryImpl) Delete(userId int) {
	var users model.User
	result := u.DB.Where("id = ?", userId).Delete(&users)
	helper.PanicIfError(result.Error)
}
