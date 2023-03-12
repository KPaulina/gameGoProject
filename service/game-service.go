package service

import "example/gameGoProject/entity"

type GameService interface {
	Save(game entity.Game) entity.Game
	FindAll() []entity.Game
}

type gameService struct {
	games []entity.Game
}

func New() GameService {
	return &gameService{}
}

func (service *gameService) Save(game entity.Game) entity.Game {
	service.games = append(service.games, game)
	return game
}

func (service *gameService) FindAll() []entity.Game {
	return service.games
}
