package model

import "time"

type Gameturn struct {
	Id				int64 		`json:"id"`
	GameId			int64		`json:"gameId"`
	UserResult		int8		`json:"userResult" validate:"min=0,max=3"`
	MachineResult 	int8		`json:"machineResult" validate:"min=0,max=3"`
	TurnType 		int8		`json:"turnType" validate:"eq=0|eq=1"`
	TurnDate 		time.Time	`json:"turnDate"`
}

func NewGameturn(id int64, gameId int64, userResult int8, machineResult int8, turnType int8, turnDate time.Time) Gameturn {
	e := Gameturn{id, gameId, userResult, machineResult, turnType, turnDate}
	return e
}
