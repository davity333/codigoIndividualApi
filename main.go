package main

import (
	dependencies "chat/Src/Endpoint/User/Infrestructure/Dependencies"
	dependenciesMessage "chat/Src/Endpoint/Message/Infrestructure/Dependencies"
	"chat/Src/Middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	sos := gin.Default()
	sos.Use(Middleware.Cors())
	dependencies.InitUserDependencies(sos)
	dependenciesMessage.InitMessageDependencies(sos)
	sos.Run(":8080")
}