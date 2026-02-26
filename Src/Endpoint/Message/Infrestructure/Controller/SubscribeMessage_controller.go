package controller

import (
	"net/http"
	"strconv"

	"chat/Src/Core/sse"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // WARNING: For development only. Secure this in production.
	},
}

type SubscribeMessageController struct {
	broadcaster *sse.Broadcaster
}

func NewSubscribeMessageController(broadcaster *sse.Broadcaster) *SubscribeMessageController {
	return &SubscribeMessageController{
		broadcaster: broadcaster,
	}
}

func (ctrl *SubscribeMessageController) Subscribe(c *gin.Context) {
	userIDStr := c.GetHeader("X-User-ID")
	if userIDStr == "" { // Attempt to get from query string if missing from header
		userIDStr = c.Query("userId")
	}

	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-ID header or userId query parameter required"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return // Ignore err, the upgrader already responded to the client
	}

	ctrl.broadcaster.Subscribe(userID, ws)

	// Listen for close/disconnection from client
	defer func() {
		ctrl.broadcaster.Unsubscribe(userID, ws)
		ws.Close()
	}()

	// The connection stays open until the client closes it.
	// We must read in a loop from the client so that we can detect
	// when the client actually disconnects from their side.
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}
