package router

import (
	"teststaticwongnai/api"

	"github.com/gin-gonic/gin"
)

func NewRouter(h api.Handler) *gin.Engine {
	r := gin.Default()
	r.GET("/getinfo", h.Getinfo)
	return r
}
