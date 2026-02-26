package controller

import (
	application "chat/Src/Endpoint/Contacts/Application"
	entities "chat/Src/Endpoint/Contacts/Domain/Entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateContactController struct {
	usecase       *application.CreateContactUseCase
	getAllUseCase *application.GetAllContactsUseCase
}

func NewCreateContactController(usecase *application.CreateContactUseCase, getAllUseCase *application.GetAllContactsUseCase) *CreateContactController {
	return &CreateContactController{
		usecase:       usecase,
		getAllUseCase: getAllUseCase,
	}
}

func (c *CreateContactController) Handle(ctx *gin.Context) {
	var req entities.Contact

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if req.UserID <= 0 || req.ContactID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "IDs inválidos"})
		return
	}

	err := c.usecase.Execute(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtener los datos completos del contacto recién creado
	contacts, err := c.getAllUseCase.Execute(req.UserID)
	if err == nil {
		for _, contact := range contacts {
			if contact.ContactID == req.ContactID {
				ctx.JSON(http.StatusCreated, gin.H{
					"message": "Contacto agregado correctamente",
					"contact": contact,
				})
				return
			}
		}
	}

	// Fallback si no se pudo obtener el contacto
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Contacto agregado correctamente",
	})
}
