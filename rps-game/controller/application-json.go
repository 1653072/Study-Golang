package controller

import "github.com/gin-gonic/gin"

func JsonController(c *gin.Context, code int, obj interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, obj)
}
