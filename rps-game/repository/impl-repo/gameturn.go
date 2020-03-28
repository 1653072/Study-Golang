package ImplRepo

import (
	"github.com/jinzhu/gorm"
	"golangdemo/rps-game/model"
	"golangdemo/rps-game/repository/interfaces"
)

type gameTurnRepo struct {
	db *gorm.DB
}

func (g gameTurnRepo) Insert(gameturn model.Gameturn) {
	g.db.Create(&gameturn)
}

func (g gameTurnRepo) FindAllGameTurnsByGameID(gameId int64) []model.Gameturn {
	var gt []model.Gameturn
	g.db.Where("game_id = ?", gameId).Find(&gt)
	return gt
}

func NewGameTurnRepository(db *gorm.DB) interfaces.GameTurnRepository {
	return &gameTurnRepo{db}
}
