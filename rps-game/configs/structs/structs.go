package structs

import "github.com/dgrijalva/jwt-go"

type MysqlConfigStruct struct {
	DbName string
	Host string
	Port string
	User string
	Password string
}

type WinRatingStruct struct {
	Username string			`json:"username"`
	WinRating float32		`json:"winRating"`
}

type JwtGenResult struct {
	Token string 			`json:"token"`
	IssuedTime int64		`json:"issuedTime"`
	ExpiredTime int64 		`json:"expiredTime"`
}

type JwtPayload struct {
	Username string `json:"username"`
}

type CustomJwtClaims struct {
	Value JwtPayload
	jwt.StandardClaims
}
