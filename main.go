package main

import (
	"go-gin-gorm-crud/config"
	"go-gin-gorm-crud/controller"
	"go-gin-gorm-crud/helper"
	"go-gin-gorm-crud/model"
	"go-gin-gorm-crud/repository"
	"go-gin-gorm-crud/router"
	"go-gin-gorm-crud/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Server Started!")

	// Config DB
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.User{})

	// Repository
	userRepository := repository.NewUsersRepositoryImpl(db)

	// Service
	usersService := service.NewUsersServiceImpl(userRepository, validate)

	// Controller
	usersController := controller.NewUsersController(usersService)

	// Router
	routes := router.NewRouter(usersController)

	server := &http.Server{
		Addr:    ":6789",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
