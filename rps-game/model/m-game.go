package model

import "time"

type Game struct {
	Id				int64 		`json:"id"`
	Username     	string	 	`json:"username" validate:"gt=6"`
	StartDate		time.Time	`json:"startDate"`
	GameResult		int8		`json:"gameResult" validate:"eq=0|eq=1"`
}

func NewGame(id int64, username string, startDate time.Time, gameResult int8) Game {
	e := Game{id, username, startDate, gameResult}
	return e
}