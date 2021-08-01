package statistics

import (
	"net/http"

	"github.com/atamano/fizz-buzz/pkg/response"
	"github.com/atamano/fizz-buzz/pkg/server"
)

type controler struct {
	service Service
}

func RegisterHandlers(routerGroup server.Router, service Service) {
	res := controler{service}

	routerGroup.GET("/stats", res.get)
}

func (r controler) get(c *server.Context) {
	stats, err := r.service.GetMostUsedRequest()
	if err != nil {
		c.JSON(http.StatusNotFound, response.BuildErrorReponse(err, "No requests found"))
		return
	}

	c.JSON(http.StatusOK, stats)
}
