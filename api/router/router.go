package router

import (
	"github.com/ady243/go-project-chat/user"
	"github.com/gin-gonic/gin"
)

var apichat *gin.Engine

func InitRoutes(userHandler *user.Handler) {
	apichat = gin.Default()

	apichat.POST("/signup", userHandler.CreateUser)

}

func Start(addr string) {
	apichat.Run(addr)
}
