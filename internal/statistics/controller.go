package statistics

import (
	"fmt"
	"net/http"

	"github.com/atamano/fizz-buzz/pkg/server"
)

type controler struct {
	service Service
}

//RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(routerGroup server.Router, service Service) {
	res := controler{service}

	fmt.Println("YES")
	routerGroup.GET("/stats", res.get)
}

func (r controler) get(c *server.Context) {

	stats, err := r.service.GetMostUsedRequest()

	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, stats)
}
