package ImplRepo

import (
	"github.com/jinzhu/gorm"
	"golangdemo/rps-game/configs/structs"
	"golangdemo/rps-game/model"
	"golangdemo/rps-game/repository/interfaces"
)

type gameRepo struct {
	db *gorm.DB
}

func (g gameRepo) Insert(game model.Game) *model.Game {
	g.db.Create(&game)
	return &game
}

func (g gameRepo) UpdateGame(game model.Game) {
	g.db.Model(&model.Game{}).Where("id = ?", game.Id).Updates(game)
}

func (g gameRepo) FindAllGamesByUsername(username string) []model.Game {
	var game []model.Game
	g.db.Where("username = ?", username).Find(&game)
	return game
}

func (g gameRepo) FindTop100WinRating() []structs.WinRatingStruct {
	var wrs []structs.WinRatingStruct
	g.db.Raw("SELECT ga.username AS username, wintable.wincount*1.0/count(*) AS win_rating FROM games ga RIGHT OUTER JOIN (SELECT g.username AS username, count(*) AS wincount FROM games g WHERE g.game_result = 1 GROUP BY g.username) AS wintable ON ga.username = wintable.username GROUP BY ga.username ORDER BY win_rating DESC LIMIT 0, 100").Scan(&wrs)
	return wrs
}

func (g gameRepo) FindGameById(gameId int64) *model.Game {
	game := &model.Game{}
	g.db.Where("id = ?", gameId).Find(&game)
	return game
}

func NewGameRepository(db *gorm.DB) interfaces.GameRepository {
	return &gameRepo{db}
}
