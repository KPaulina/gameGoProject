package main

import (
	"example/gameGoProject/service"
	"github.com/gin-gonic/gin"
)

//	func init() {
//		initializers.LoadEnvVariables()
//	}
var (
	gameService     service.GameService       = service.New()
	videoController controller.GameController = service.New(GameService)
)

func main() {
	server.GET("/games", func(ctx *gin.Context) {
		ctx.JSON(200, gameController.FindAll())
	})
	server.POST("/games", func(ctx *gin.Context) {
		ctx.JSON(200, gameController.Save(ctx))
	})

	server.Run(":8080")
}
