package dependencies

import (
	application "chat/Src/Endpoint/User/Application"
	controller "chat/Src/Endpoint/User/Infrestructure/Controller"
	router "chat/Src/Endpoint/User/Infrestructure/Router"
	sql "chat/Src/Endpoint/User/Infrestructure/Sql"

	"github.com/gin-gonic/gin"
)

func InitUserDependencies(e *gin.Engine) {
	ps, err := sql.NewMySQL()
	if err != nil {
		panic(err)
	}

	getAllProductUseCase := application.NewGetAllUsersUseCase(ps)
	getAllProductController := controller.NewGetAllUserController(getAllProductUseCase)

	getUserByNameUseCase := application.NewGetUserByNameUseCase(ps)
	getUserByNameController := controller.NewGetUserByNameController(getUserByNameUseCase)

	createUserUseCase := application.NewCreateUserUseCase(ps)
	createUserController := controller.NewCreateUserController(createUserUseCase)

	loginUserUseCase := application.NewLoginUserUseCase(ps)
	loginUserController := controller.NewLoginUserUseCase(loginUserUseCase)
	
	router.RouterUser(e, getAllProductController, getUserByNameController, createUserController, loginUserController)
}