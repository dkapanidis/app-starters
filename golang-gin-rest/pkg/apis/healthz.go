package apis

import (
	"net/http"

	"github.com/dkapanidis/app-starters/golang-gin-rest/pkg/store"
	"github.com/gin-gonic/gin"
)

// Healthz handles healthz
// @Summary healthz
// @Description healthz
// @Tags healthz
// @Accept json
// @Produce json
// @Success 204 {string} string	"ok"
// @Router /healthz [get]
func Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, store.Healthz())
}
