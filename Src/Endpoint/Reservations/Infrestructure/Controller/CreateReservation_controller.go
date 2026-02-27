package controller

import (
	application "chat/Src/Endpoint/Reservations/Application"
	entities "chat/Src/Endpoint/Reservations/Domain/Entities"
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

type CreateReservationController struct {
	usecase *application.CreateReservationUseCase
}

func NewCreateReservationController(usecase *application.CreateReservationUseCase) *CreateReservationController {
	return &CreateReservationController{
		usecase: usecase,
	}
}

func (ctrl *CreateReservationController) CreateReservation(c *gin.Context) {

    // 1. Leer body crudo
    bodyBytes, _ := io.ReadAll(c.Request.Body)
    fmt.Println("📩 BODY RECIBIDO:", string(bodyBytes))

    // 2. Reponer body para BindJSON
    c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

    // 3. Parsear JSON
    var req entities.Reservation
    if err := c.BindJSON(&req); err != nil {
        fmt.Println("❌ ERROR PARSEANDO JSON:", err)
        c.JSON(400, gin.H{"error": "JSON inválido", "detail": err.Error()})
        return
    }

    // 4. Log del objeto parseado
    fmt.Printf("📦 JSON PARSEADO: %+v\n", req)

    // 5. Lógica normal
    if err := ctrl.usecase.Execute(&req); err != nil {
        c.JSON(500, gin.H{"error": "Error interno del servidor", "detail": err.Error()})
        return
    }

    c.JSON(201, gin.H{"message": "Reservación creada", "data": req})
}

