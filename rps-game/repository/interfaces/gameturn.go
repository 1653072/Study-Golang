package interfaces

import (
	"golangdemo/rps-game/model"
)

type GameTurnRepository interface {
	Insert(gameturn model.Gameturn)
	FindAllGameTurnsByGameID(gameId int64) []model.Gameturn
}
