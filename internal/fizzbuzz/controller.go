package fizzbuzz

import (
	"errors"
	"net/http"

	"fizzbuzz/internal/statistics"
	"fizzbuzz/pkg/response"
	"fizzbuzz/pkg/server"
)

type controler struct {
	service      Service
	statsService statistics.Service
}

func RegisterHandlers(routerGroup server.Router, service Service, statsService statistics.Service) {
	res := controler{service, statsService}

	routerGroup.POST("/fizzbuzz", res.post)
}

func (r controler) post(c *server.Context) {
	var params postRequest

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, response.BuildErrorReponse(err, "Error validating parameters"))
		return
	}

	if params.Int1 > params.Limit || params.Int2 > params.Limit {
		c.JSON(http.StatusBadRequest, response.BuildErrorReponse(errors.New("Bad limit"), "Error validating parameters"))
		return
	}

	r.statsService.IncrementRequestCount(params)
	res := r.service.Compute(params)

	c.JSON(http.StatusOK, struct {
		Result string `json:"result"`
	}{res})
}
