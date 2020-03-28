package routers

import (
	"github.com/gin-gonic/gin"
	"golangdemo/rps-game/configs/api"
	LogConf "golangdemo/rps-game/configs/log-conf"
	"golangdemo/rps-game/configs/system-code"
	SystemPath "golangdemo/rps-game/configs/system-path"
	"golangdemo/rps-game/controller"
	"golangdemo/rps-game/helpers/logging"
	"golangdemo/rps-game/middlewares"
)

func InitializeRouters() error {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CorsMiddleware())

	// Login router
	r.POST(api.LoginUrl, controller.LoginAccount)

	// Account registration router
	r.POST(api.RegisterAccountUrl, controller.CreateAccount)

	// Game group router
	gameUrlGroup := r.Group(api.GameBaseUrl)
	gameUrlGroup.Use(middlewares.JwtAuthMiddleware)
	{
		gameUrlGroup.POST(api.PlayGameUrl, controller.PlayGame)
		gameUrlGroup.GET(api.SeeGameHistoryUrl, controller.GetGameHistory)
		gameUrlGroup.GET(api.SeeTopPlayersUrl, controller.GetTopPlayers)
		gameUrlGroup.GET(api.SeeGameTurnsUrl, controller.GetGameTurnsHistory)
	}

	logging.SysLog.Println(logging.FormatResult(LogConf.Info, nil, "Server listening on " + api.DomainUrl))
	rErr := r.RunTLS(api.DomainUrl, SystemPath.SCertFilePath, SystemPath.SKeyFilePath)
	if rErr != nil {
		logging.SysLog.Printf(logging.FormatResult(LogConf.Fatal, SystemCode.ErrRunServer, rErr.Error()))
		return SystemCode.ErrRunServer
	}

	return nil
}