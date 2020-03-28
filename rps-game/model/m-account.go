package model

type Account struct {
	Username     string	 		`json:"username" validate:"gte=6" binding:"required"`
	Password     string  		`json:"password" validate:"gte=6" binding:"required"`
}

func NewAccount(username, password string) Account {
	e := Account{Username: username, Password: password}
	return e
}
