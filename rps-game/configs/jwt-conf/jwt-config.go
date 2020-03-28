package JwtConf

import "time"

var (
	JwtCookieName = "JwtToken"
	// In JWT, the expiry time is expressed as unix milliseconds
	JwtExpirationTime = time.Now().Add(168 * time.Hour).Unix()
)
