package controller

import (
	"net/http"
	"strconv"

	"chat/Src/Core/sse"

	"github.com/gin-gonic/gin"
)

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
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-ID header required"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid X-User-ID"})
		return
	}

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("X-Accel-Buffering", "no")

	messages := ctrl.broadcaster.Subscribe(userID)
	defer ctrl.broadcaster.Unsubscribe(userID, messages)

	c.Header("X-Message", "Connected")

	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Streaming not supported"})
		return
	}

	for {
		select {
		case msg, ok := <-messages:
			if !ok {
				return
			}
			c.SSEvent("message", msg)
			flusher.Flush()
		case <-c.Request.Context().Done():
			return
		}
	}
}
