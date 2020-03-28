package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	LogConf "golangdemo/rps-game/configs/log-conf"
	SystemCode "golangdemo/rps-game/configs/system-code"
	JwtHelper "golangdemo/rps-game/helpers/jwt"
	"golangdemo/rps-game/helpers/logging"
	"golangdemo/rps-game/model"
	"golangdemo/rps-game/repository"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var drawTurnMap = sync.Map{}

func playOfMachine() int8 {
	//0: Kéo (Scissor)
	//1: Búa (Rock)
	//2: Bao (Paper)

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 2
	return int8(rand.Intn(max-min+1) + min)
}

func checkGameResult(userResult int8, machineResult int8) int8 {
	/*
	 * Return 0: Hòa (No one wins)
	 * Return 1: User wins
	 * Return -1: User loses
	 */

	if userResult == machineResult {
		return 0
	}

	switch userResult {
	case 0: //user: Kéo
		{
			if machineResult == 1 {
				return -1
			}
			return 1
		}
	case 1: //user: Búa
		{
			if machineResult == 2 {
				return -1
			}
			return 1
		}
	default: //user: Bao
		{
			if machineResult == 0 {
				return -1
			}
			return 1
		}
	}
}

func GetGameHistory(c *gin.Context) {
	// Get data from Jwt
	dataJwt, dErr := JwtHelper.GetPayloadFromJwtWithGin(c)
	if len(dErr) > 0 {
		logging.SysLog.Println(logging.FormatResult(LogConf.Warn, nil, "[GetGameHistory] Get username from Jwt of user fail"))
		JsonController(c, http.StatusBadRequest, gin.H{
			"message": dErr,
		})
		return
	}

	username := dataJwt.Username

	// Get game history
	gameList := repository.GameRepo.FindAllGamesByUsername(username)

	// Done
	JsonController(c, http.StatusOK, gin.H{
		"gameHistory": gameList,
		"message": SystemCode.GetGameHistorySuccessfully,
	})
}

func GetGameTurnsHistory(c *gin.Context) {
	// Query string parameters
	gameId := c.Query("gameid")

	if len(gameId) <= 0 {
		logging.SysLog.Println(logging.FormatResult(LogConf.Warn, nil, "[GetGameTurnsHistory] Lack of field in request param"))
		JsonController(c, http.StatusBadRequest, gin.H{
			"message": SystemCode.LackOfFieldInRequestParam,
		})
		return
	}

	// Convert string to int64
	gId, pErr := strconv.ParseInt(gameId, 10, 64)
	if pErr != nil {
		logging.SysLog.Println(logging.FormatResult(LogConf.Warn, nil, "[GetGameTurnsHistory] Parse string to int64 fail"))
		JsonController(c, http.StatusServiceUnavailable, gin.H{
			"message": SystemCode.SomethingWentWrong,
		})
		return
	}

	// Get game history
	gameTurnList := repository.GameTurnRepo.FindAllGameTurnsByGameID(gId)

	// Done
	JsonController(c, http.StatusOK, gin.H{
		"gameId": gId,
		"gameTurnHistory": gameTurnList,
		"message": SystemCode.GetGameTurnHistorySuccessfully,
	})
}

func GetTopPlayers(c *gin.Context) {
	// Get top players
	winRatingList := repository.GameRepo.FindTop100WinRating()

	// Done
	JsonController(c, http.StatusOK, gin.H{
		"topPlayers": winRatingList,
		"message": SystemCode.GetGameTopPlayersSuccessfully,
	})
}

func PlayGame(c *gin.Context) {
	// Get data from Jwt
	dataJwt, dErr := JwtHelper.GetPayloadFromJwtWithGin(c)
	if len(dErr) > 0 {
		logging.SysLog.Println(logging.FormatResult(LogConf.Warn, nil, "[PlayGame] Get username from Jwt of user fail"))
		JsonController(c, http.StatusBadRequest, gin.H{
			"message": dErr,
		})
		return
	}

	username := dataJwt.Username

	// Get data from request body & Checking with conditions
	bodyBuf, err := c.GetRawData()
	if err != nil {
		logging.SysLog.Println(logging.FormatResult(LogConf.Warn, nil, "[PlayGame] Request body is inappropriate"))
		JsonController(c, http.StatusBadRequest, gin.H{
			"message": SystemCode.InappropriateRequestBody,
		})
		return
	}

	var bodyData map[string]interface{}
	if err := json.Unmarshal(bodyBuf, &bodyData); err != nil {
		logging.SysLog.Println(logging.FormatResult(LogConf.Warn, nil, "[PlayGame] Request body is inappropriate"))
		JsonController(c, http.StatusBadRequest, gin.H{
			"message": SystemCode.InappropriateRequestBody,
		})
		return
	}

	var userChoice int8
	if bodyData["choice"] == nil {
		logging.SysLog.Println(logging.FormatResult(LogConf.Warn, nil, "[PlayGame] Lack of field `choice` in request body"))
		JsonController(c, http.StatusBadRequest, gin.H{
			"message": SystemCode.LackOfFieldInRequestBody + " with 'choice' field",
		})
		return
	} else {
		userChoice = int8(bodyData["choice"].(float64))
		if userChoice < 0 || userChoice > 2 {
			logging.SysLog.Println(logging.FormatResult(LogConf.Warn, nil, "[PlayGame] The `choice` value is not in suitable range"))
			JsonController(c, http.StatusBadRequest, gin.H{
				"message": "The value of 'choice' field must be in {0 (Scissor), 1 (Rock), 2 (Paper)}",
			})
			return
		}
	}

	// Play game
	machineChoice := playOfMachine()
	gameResult := checkGameResult(userChoice, machineChoice)

	// Check game status
	/*
		+ The "gameResult" has 3 values, it can equal 1 (User wins the game), 0 (User draws the game) or
		-1 (User loses the game).
		+ If the "status" is false, that means the new game will be created. Otherwise, no new-game is
		created, only new-game-turn is created.
		+ In addition, when "status" is false, "gameResult" will be changed to "gr". The "gr" only has 2
		values, it can be 1 (User wins the game) or 0 (User loses or draws the game). The situation (draw)
		will be solved by the "gameId", which is added to the "drawTurnMap" for the next game turns. The
		next game turns of user will be based on the previous "gameId" value, and they also determine final
		"gameResult" of the user that it will be changed from 0 to 1 (if user wins) or keep 0 (if user loses).
	*/
	var gameId int64
	_, status := drawTurnMap.Load(username)
	if !status {
		// Create new game
		var gr int8
		if gameResult == 1 {
			gr = 1
		} else {
			gr = 0
		}
		newGame := model.Game{Username: username, GameResult: gr, StartDate: time.Now()}
		insertedGame := repository.GameRepo.Insert(newGame)

		gameId = insertedGame.Id

		// If no one wins/loses => drawTurnMap put username & gameId which will be used for next game turn
		if gameResult == 0 {
			drawTurnMap.Store(username, gameId)
		}
	} else {
		// Last game has noone wins/loses => New turn of current user is needed
		val, _ := drawTurnMap.Load(username)
		gameId = val.(int64)
		if gameResult == -1 || gameResult == 1 {
			drawTurnMap.Delete(username)
			if gameResult == 1 {
				newGame := model.Game{Id: gameId, GameResult: gameResult}
				repository.GameRepo.UpdateGame(newGame)
			}
		}
	}

	// Type of game turn: 0 (Finished game) | 1 (Additional game turn)
	var turnType int8
	if gameResult != 0 {
		turnType = 0
	} else {
		turnType = 1
	}

	// Add new game turn to DB
	newGameTurn := model.Gameturn{GameId: gameId, MachineResult: machineChoice, UserResult: userChoice, TurnType: turnType, TurnDate: time.Now()}
	repository.GameTurnRepo.Insert(newGameTurn)

	// Done
	var messageResponse = ""
	switch(gameResult) {
		case -1:
			messageResponse = "So sorry! You lose the game."
			break
		case 0:
			messageResponse = "This turn is draw, try more."
			break
		case 1:
			messageResponse = "Congratulations! You won the game."
			break;
	}

	JsonController(c, http.StatusOK, gin.H{
		"gameId": gameId,
		"machineChoice": machineChoice,
		"yourChoice": userChoice,
		"message": messageResponse,
	})
}
