package fizzbuzz

import (
	"net/http"

	"github.com/atamano/fizz-buzz/internal/statistics"
	"github.com/gin-gonic/gin"
)

type controler struct {
	service      Service
	statsService statistics.Service
}

//RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(routerGroup *gin.RouterGroup, service Service, statsService statistics.Service) {
	res := controler{service, statsService}

	routerGroup.POST("/fizzbuzz", res.post)
}

func (r controler) post(c *gin.Context) {
	var params postRequest

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := params.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Could be handled with pub/sub workers on larger app
	r.statsService.IncrementRequestCount(params)

	res := r.service.Compute(params)
	c.JSON(http.StatusOK, gin.H{"result": res})
}
