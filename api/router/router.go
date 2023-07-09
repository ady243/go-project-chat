package router

import (
	"github.com/ady243/go-project-chat/internal/user"
	websocket "github.com/ady243/go-project-chat/internal/webSocket"
	"github.com/gin-gonic/gin"
)

var apichat *gin.Engine

func InitRoutes(userHandler *user.Handler, wsHandler *websocket.Handler) {
	apichat = gin.Default()

	apichat.POST("/signup", userHandler.CreateUser)
	apichat.POST("/login", userHandler.Login)
	apichat.POST("/logout", userHandler.Logout)
	apichat.DELETE("/users/:id", userHandler.DeleteUser)

	apichat.POST("/websocket/createRoom", wsHandler.CreateRoom)
	apichat.GET("/websocket/joinRoom/:roomId", wsHandler.JoinRoom)
	apichat.GET("/websocket/getRooms", wsHandler.GetRooms)
	apichat.GET("websocket/getClients/:roomId", wsHandler.GetClients)

}

func Start(addr string) {
	apichat.Run(addr)
}
