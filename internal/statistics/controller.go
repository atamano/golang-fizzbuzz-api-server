package statistics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type controler struct {
	service Service
}

//RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(routerGroup *gin.RouterGroup, service Service) {
	res := controler{service}

	routerGroup.GET("/stats", res.get)
}

func (r controler) get(c *gin.Context) {

	stats, err := r.service.GetMostUsedRequest()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, stats)
}
