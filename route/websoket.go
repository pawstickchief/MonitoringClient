package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
)

// 初始化 WebSocket 升级器
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有请求，生产环境中应根据需要进行限制
		return true
	},
}

// WebSocket 处理器
func WebsocketHandler(c *gin.Context) {
	// 从 URL 查询参数中获取用户ID（或其他标识）
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// 升级 HTTP 连接到 WebSocket 连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.L().Info("Failed to upgrade connection: %v", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade to websocket"})
		return
	}
	defer conn.Close()
	zap.L().Info("WebSocket connection established for user:",
		zap.String("userID", userID),
		zap.Error(err),
	)

	for {
		// 读取消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			zap.L().Info("Read error for user",
				zap.String("userID", userID),
				zap.Error(err),
			)
			break
		}

		zap.L().Info("Received message from user",
			zap.String("userID", userID),
			zap.String("message", string(message)),
			zap.Error(err),
		)
		// 发送回消息
		responseMessage := fmt.Sprintf("User %s: %s", userID, string(message))
		err = conn.WriteMessage(messageType, []byte(responseMessage))
		if err != nil {
			zap.L().Info("Write error for user",
				zap.String("userID", userID),
				zap.Error(err),
			)
			break
		}
	}
}
