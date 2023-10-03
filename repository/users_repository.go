package repository

import "go-gin-gorm-crud/model"

type UsersRepository interface {
	Create(user model.User)
	FindAll() []model.User
	FindById(userId int) (user model.User, err error)
	Update(user model.User)
	Delete(userId int)
}
