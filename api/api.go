package api

import (
	"net/http"
	"teststaticwongnai/duplicate"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=./api.go -destination=./mock_api/mock_api.go -package=mock_api
type TestGetinfo interface {
	ResGet() (map[string]int, duplicate.Age)
}

type Handler struct {
	ServiceRes TestGetinfo
}

func (h Handler) Getinfo(c *gin.Context) {
	Dup_map, Dup_Age := h.ServiceRes.ResGet()

	c.JSON(http.StatusOK, gin.H{"Province": Dup_map, "AgeGroup": Dup_Age})

}
