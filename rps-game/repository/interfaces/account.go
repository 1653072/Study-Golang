package interfaces

import (
	"golangdemo/rps-game/model"
)

type AccountRepository interface {
	UpdateAccount(account model.Account)
	Insert(account model.Account) (*model.Account, bool)
	FindByUsername(username string) *model.Account
}
