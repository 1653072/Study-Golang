package main

import (
	"golangdemo/rps-game/configs/log-conf"
	"golangdemo/rps-game/configs/system-code"
	_ "golangdemo/rps-game/configs/system-path"
	_ "golangdemo/rps-game/helpers/jwt"
	"golangdemo/rps-game/helpers/logging"
	"golangdemo/rps-game/model"
	"golangdemo/rps-game/repository"
	"golangdemo/rps-game/routers"
)

func main() {
	// Initialize logging mechanism
	logging.InitializeLogging()

	// Initialize model
	modelError := model.InitializeModel()
	if modelError != nil {
		logging.SysLog.Println(logging.FormatResult(LogConf.Fatal, SystemCode.ErrRunServer, "Due to error, app stopped"))
		return
	}
	defer model.CloseModel()

	// Initialize repositories for uses
	repository.InitializeRepositories(model.MysqlInstance)

	// Initialize router
	routerError := routers.InitializeRouters()
	if routerError != nil {
		logging.SysLog.Println(logging.FormatResult(LogConf.Fatal, SystemCode.ErrRunServer, "Due to error, app stopped"))
		return
	}
}
