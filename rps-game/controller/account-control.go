package controller

import (
	"github.com/gin-gonic/gin"
	JwtConfig "golangdemo/rps-game/configs/jwt-conf"
	LogConf "golangdemo/rps-game/configs/log-conf"
	"golangdemo/rps-game/configs/system-code"
	JwtHelper "golangdemo/rps-game/helpers/jwt"
	"golangdemo/rps-game/helpers/logging"
	"golangdemo/rps-game/helpers/utils"
	"golangdemo/rps-game/helpers/validator"
	"golangdemo/rps-game/model"
	"golangdemo/rps-game/repository"
	"net/http"
)

func LoginAccount(c *gin.Context) {
	// Validate user account
	var account model.Account
	_ = c.BindJSON(&account)

	vErr := validator.ValidateAccountModel(account)
	if !vErr {
		logging.SysLog.Println(logging.FormatResult(LogConf.Warn, nil, "[LoginAccount] Validate account with username " + account.Username + " fail"))
		JsonController(c, http.StatusBadRequest, gin.H{
			"message": SystemCode.InappropriateAccountInfo,
		})
		return
	}

	// Get account from database
	dbAccount := repository.AccountRepo.FindByUsername(account.Username)
	if dbAccount == nil {
		logging.SysLog.Println(logging.FormatResult(LogConf.Warn, nil, "[LoginAccount] Account with username " + account.Username + " doesn't exist"))
		JsonController(c, http.StatusBadRequest, gin.H{
			"message": SystemCode.InappropriateAccountInfo,
		})
		return
	}

	// Compare password
	cRes := utils.CompareBcryptPassword(dbAccount.Password, []byte(account.Password))
	if !cRes {
		logging.SysLog.Println(logging.FormatResult(LogConf.Warn, nil, "[LoginAccount] Username " + account.Username + " login with wrong password"))
		JsonController(c, http.StatusBadRequest, gin.H{
			"message": SystemCode.InappropriateAccountInfo,
		})
		return
	}

	// Generate token
	jwtResult, jErr := JwtHelper.GenerateJwt(account.Username)
	if jErr != nil {
		// Inside "GenerateJwt" function has log, no need log here
		JsonController(c, http.StatusServiceUnavailable, gin.H{
			"message": SystemCode.SomethingWentWrong,
		})
		return
	}

	// Set cookie
	maxAge := int(jwtResult.ExpiredTime - jwtResult.IssuedTime)
	c.SetCookie(JwtConfig.JwtCookieName, jwtResult.Token, maxAge, "/", "", http.SameSiteLaxMode, true, true)

	// Done
	JsonController(c, http.StatusOK, gin.H{
		"token": jwtResult.Token,
		"expiredTime": jwtResult.ExpiredTime,
		"message": SystemCode.LoginInSuccessfully,
	})
}

func CreateAccount(c *gin.Context) {
	// Validate user account
	var account model.Account
	_ = c.BindJSON(&account)

	vErr := validator.ValidateAccountModel(account)
	if !vErr {
		logging.SysLog.Println(logging.FormatResult(LogConf.Warn, nil, "[CreateAccount] Validate account with username " + account.Username + " fail"))
		JsonController(c, http.StatusBadRequest, gin.H{
			"message": SystemCode.InappropriateAccountInfo,
		})
		return
	}

	// Get bcrypt password of account
	account.Password = utils.GenerateBcryptPassword([]byte(account.Password))

	_, resStatus := repository.AccountRepo.Insert(account)
	if !resStatus {
		logging.SysLog.Println(logging.FormatResult(LogConf.Warn, nil, "[CreateAccount] Account with username " + account.Username + " existed, create the new one fail"))
		JsonController(c, http.StatusBadRequest, gin.H{
			"message": SystemCode.AccountExisted,
		})
		return
	}

	// Done
	JsonController(c, http.StatusOK, gin.H{
		"message": SystemCode.RegisterAccountSuccessfully,
	})
}
