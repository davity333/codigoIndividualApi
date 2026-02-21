package dependencies

import (
	application "chat/Src/Endpoint/Message/Application"
	controller "chat/Src/Endpoint/Message/Infrestructure/Controller"
	router "chat/Src/Endpoint/Message/Infrestructure/Router"
	sql "chat/Src/Endpoint/Message/Infrestructure/Sql"
	"github.com/gin-gonic/gin"
)

func InitMessageDependencies(gin *gin.Engine) {
	ps, err := sql.NewMySQL()
	if err != nil {
		panic(err)
	}

	getAllMessageByUserIdUseCase := application.NewGetMessageByUserIdUseCase(ps)
	getAllMessageByUserController := controller.NewGetAllMessageController(getAllMessageByUserIdUseCase)

	sendMessageByUserUseCase := application.NewSendMessageUseCase(ps)
	sendMessageByUserController := controller.NewSendMessageController(*sendMessageByUserUseCase)

	deleteMessageByUserUseCase := application.NewDeleteMessageUseCase(ps)
	deleteMessageByUserController := controller.NewDeleteMessageController(*deleteMessageByUserUseCase)

	router.MessageRouter(gin, getAllMessageByUserController, sendMessageByUserController, deleteMessageByUserController)
}