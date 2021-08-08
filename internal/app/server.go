package app

import "github.com/gin-gonic/gin"

func (a *App) NewServer() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", a.ping)

	return r
}
