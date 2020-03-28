package repository

import (
	"github.com/jinzhu/gorm"
	ImplRepo "golangdemo/rps-game/repository/impl-repo"
	"golangdemo/rps-game/repository/interfaces"
)

var (
	AccountRepo			interfaces.AccountRepository 	= nil
	GameRepo			interfaces.GameRepository		= nil
	GameTurnRepo		interfaces.GameTurnRepository	= nil
)

func InitializeRepositories(db *gorm.DB) {
	if AccountRepo == nil && GameRepo == nil && GameTurnRepo == nil {
		AccountRepo = ImplRepo.NewAccountRepository(db)
		GameRepo = ImplRepo.NewGameRepository(db)
		GameTurnRepo = ImplRepo.NewGameTurnRepository(db)
	}
}
