package router

import (
	"go-gin-gorm-crud/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(usersController *controller.UsersController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Server Online!")
	})

	baseRouter := router.Group("/api")
	userRouter := baseRouter.Group("/users")
	userRouter.GET("", usersController.FindAll)
	userRouter.GET("/:userId", usersController.FindById)
	userRouter.POST("", usersController.Create)
	// userRouter.PUT("/:userId", usersController.Update)
	// userRouter.DELETE("/:userId", usersController.Delete)

	return router;
}
