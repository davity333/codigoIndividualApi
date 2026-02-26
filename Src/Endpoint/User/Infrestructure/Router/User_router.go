package router

import (
	controller "chat/Src/Endpoint/User/Infrestructure/Controller"

	"github.com/gin-gonic/gin"
)

func RouterUser(g *gin.Engine,
	getAllUserController *controller.GetAllUserController,
	getUserByNameController *controller.GetUserByNameController,
	createUserController *controller.CreateUserController,
	loginUserController *controller.LoginUserUseCase,
	GetTeacherByIDController *controller.GetTeacherByIDController,
) {
	userGroup := g.Group("/api/v1/users")
	{
		userGroup.GET("/getAll", getAllUserController.GetUser)
		userGroup.GET("/:username", getUserByNameController.GetByUsername)
		userGroup.POST("/create", createUserController.CreateUser)
		userGroup.POST("/login", loginUserController.LoginUser)
		userGroup.GET("/teacher/:id", GetTeacherByIDController.Handle)
	}
}
