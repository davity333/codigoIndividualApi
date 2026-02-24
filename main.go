package main

import (
	coredi "chat/Src/Core/di"
	"chat/Src/Middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	sos := gin.Default()
	sos.Use(Middleware.Cors())
	if err := coredi.InitRoutes(sos); err != nil {
		panic(err)
	}
	sos.Run(":8080")
}
