package middlewares

import (
	"github.com/gin-gonic/gin"
	JwtConfig "golangdemo/rps-game/configs/jwt-conf"
	SystemCode "golangdemo/rps-game/configs/system-code"
	JwtHelper "golangdemo/rps-game/helpers/jwt"
	"net/http"
)

func JwtAuthMiddleware(c *gin.Context) {
	tokenCookie, tcErr := c.Cookie(JwtConfig.JwtCookieName)
	if tcErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": SystemCode.JwtTokenNotFound,
		})
		return
	}

	res, err := JwtHelper.VerifyJwt(tokenCookie);
	if !res {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
}
