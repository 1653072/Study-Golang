package JwtHelper

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	JwtConfig "golangdemo/rps-game/configs/jwt-conf"
	LogConf "golangdemo/rps-game/configs/log-conf"
	"golangdemo/rps-game/configs/structs"
	SystemCode "golangdemo/rps-game/configs/system-code"
	SystemPath "golangdemo/rps-game/configs/system-path"
	"golangdemo/rps-game/helpers/logging"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var jwtKey = readJwtKeyFromFile()

func readJwtKeyFromFile() []byte {
	file, err := os.Open(SystemPath.JWTKeyFilePath)
	if err != nil {
		logging.SysLog.Fatal(logging.FormatResult(LogConf.Fatal, SystemCode.ErrLoadFileFail, err.Error()))
	}
	defer file.Close()

	content, _ := ioutil.ReadAll(file)
	return content
}

func getTokenFromAuth(tok string) string {
	if len(tok) <= 0 {
		return ""
	}
	s := strings.Split(tok, " ")
	return s[1]
}

func GenerateJwt(username string) (*structs.JwtGenResult, error) {
	var expTime = JwtConfig.JwtExpirationTime
	var issuedAt = time.Now().Unix()

	myContentClaim := structs.JwtPayload{
		Username: username,
	}

	myClaim := structs.CustomJwtClaims{
		Value: myContentClaim,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime, // exp
			IssuedAt: issuedAt, // iat
		},
	}

	readyToken := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)
	finalToken, tErr := readyToken.SignedString(jwtKey)
	if tErr != nil {
		logging.SysLog.Println(logging.FormatResult(LogConf.Error, SystemCode.ErrGenerateJwtFail, tErr.Error()))
		return nil, SystemCode.ErrGenerateJwtFail
	}

	finalToken = "Bearer " + finalToken
	result := structs.JwtGenResult{Token: finalToken, IssuedTime: issuedAt, ExpiredTime: expTime}
	return &result, nil
}

func GetJwtClaims(tok string) (*structs.CustomJwtClaims, string) {
	userToken := getTokenFromAuth(tok)

	token, err := jwt.ParseWithClaims(userToken, &structs.CustomJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, SystemCode.InappropriateJwtToken
	}

	if claims, ok := token.Claims.(*structs.CustomJwtClaims); ok && token.Valid {
		return claims, ""
	} else {
		return nil, SystemCode.InappropriateJwtToken
	}
}

func VerifyJwt(tok string) (bool, string) {
	userToken := getTokenFromAuth(tok)

	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if token == nil {
		return false, SystemCode.InappropriateJwtToken
	}

	if token.Valid {
		return true, ""
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, SystemCode.InappropriateJwtToken
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, SystemCode.ExpiredJwtToken
		} else {
			return false, SystemCode.InappropriateJwtToken
		}
	} else {
		return false, SystemCode.InappropriateJwtToken
	}
}

func GetPayloadFromJwtWithGin(c *gin.Context) (*structs.JwtPayload, string) {
	// Get token from Cookie => Check Jwt & Get data field
	tokenCookie, tcErr := c.Cookie(JwtConfig.JwtCookieName)
	if tcErr != nil {
		return nil, SystemCode.JwtTokenNotFound
	}

	claims, cErr := GetJwtClaims(tokenCookie)
	if len(cErr) > 0 {
		return nil, SystemCode.InappropriateJwtToken
	}

	return &claims.Value, ""
}