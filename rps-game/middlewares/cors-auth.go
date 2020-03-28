package middlewares

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"net/http"
)

func CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"https://foo.bar"},
		AllowCredentials: true,
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowedHeaders: []string{"Content-Type", "Content-Length", "Accept", "Accept-Encoding", "Authorization", "Origin"},
		//Enable Debugging for testing, consider disabling in production
		//Debug: true,
	})
}
