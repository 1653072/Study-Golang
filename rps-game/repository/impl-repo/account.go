package ImplRepo

import (
	"github.com/jinzhu/gorm"
	"golangdemo/rps-game/model"
	"golangdemo/rps-game/repository/interfaces"
)

type accRepo struct {
	db *gorm.DB
}

func (a accRepo) UpdateAccount(account model.Account) {
	a.db.Model(&model.Account{}).Where("username = ?", account.Username).Updates(account)
}

func (a accRepo) Insert(account model.Account) (*model.Account, bool) {
	isNotFound := a.db.First(&model.Account{}, "username = ?", account.Username).RecordNotFound()
	if !isNotFound {
		return nil, false
	}
	a.db.Create(&account)
	return &account, true
}

func (a accRepo) FindByUsername(username string) *model.Account {
	account := model.Account{}
	err := a.db.Where("username = ?", username).Find(&account).RecordNotFound()
	if err {
		return nil
	}
	return &account
}

func NewAccountRepository(db *gorm.DB) interfaces.AccountRepository {
	return &accRepo{db}
}
