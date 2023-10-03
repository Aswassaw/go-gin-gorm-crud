package controller

import (
	"go-gin-gorm-crud/data/request"
	"go-gin-gorm-crud/data/response"
	"go-gin-gorm-crud/helper"
	"go-gin-gorm-crud/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersService service.UsersService
}

func NewUsersController(service service.UsersService) *UsersController {
	return &UsersController{
		usersService: service,
	}
}

// Create Controller
func (controller *UsersController) Create(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.PanicIfError(err)

	controller.usersService.Create(createUserRequest)
	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll Controller
func (controller *UsersController) FindAll(ctx *gin.Context) {
	usersResponse := controller.usersService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   usersResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindById Controller
func (controller *UsersController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.usersService.FindById(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Update Controller
func (controller *UsersController) Update(ctx *gin.Context) {
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	helper.PanicIfError(err)

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)
	updateUserRequest.Id = id

	controller.usersService.Update(updateUserRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Delete Controller
func (controller *UsersController) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.usersService.Delete(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
