package router

import (
	controller "chat/Src/Endpoint/Contacts/Infrestructure/Controller"

	"github.com/gin-gonic/gin"
)

func RouterContact(g *gin.Engine,
	getAllContactsController *controller.GetAllContactsController,
	createContactController *controller.CreateContactController,
	deleteContactController *controller.DeleteContactController,
	getContactByNameController *controller.GetContactByNameController,
) {
	contactGroup := g.Group("/api/v1/contacts")
	{
		contactGroup.GET("/getAll/:userId", getAllContactsController.Handle)
		contactGroup.POST("/create", createContactController.Handle)
		contactGroup.DELETE("/delete/:userId/:contactId", deleteContactController.Handle)
		contactGroup.GET("/search/:username", getContactByNameController.Handle)
	}
}
