package validator

import (
	"github.com/go-playground/validator"
	"golangdemo/rps-game/model"
)

var validate = validator.New()

func ValidateAccountModel(ac model.Account) bool {
	err := validate.Struct(ac)
	if err != nil {
		return false
	}
	return true
}

func ValidateGameModel(g model.Game) bool {
	err := validate.Struct(g)
	if err != nil {
		return false
	}
	return true
}

func ValidateGameTurnModel(gt model.Gameturn) bool {
	err := validate.Struct(gt)
	if err != nil {
		return false
	}
	return true
}