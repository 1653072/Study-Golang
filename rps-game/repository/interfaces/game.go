package interfaces

import (
	"golangdemo/rps-game/configs/structs"
	"golangdemo/rps-game/model"
)

type GameRepository interface {
	Insert(game model.Game) *model.Game
	UpdateGame(game model.Game)
	FindAllGamesByUsername(username string) []model.Game
	FindTop100WinRating() []structs.WinRatingStruct
	FindGameById(gameId int64) *model.Game
}
