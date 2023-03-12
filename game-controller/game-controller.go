package game_controller

import (
	"example/gameGoProject/entity"
	"example/gameGoProject/service"
	"github.com/gin-gonic/gin"
)

type GameController interface {
	FindAll() []entity.Game
	Save(ctx *gin.Context)
}

type controller struct {
	service service.GameService
}

func New(service service.GameService) GameController {
	return controller{
		service: service,
	}
}

func (c *controller) FindAll []entity.Game {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Game {
	var game entity.Game
	ctx.BindJSON(&game)
	c.service.Save(game)
	return game
}