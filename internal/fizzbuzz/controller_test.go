package fizzbuzz

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atamano/fizz-buzz/internal/statistics"
	"github.com/atamano/fizz-buzz/pkg/logger"
	"github.com/atamano/fizz-buzz/pkg/server"
	"github.com/stretchr/testify/assert"
)

type testCasesController struct {
	testCase   string
	request    postRequest
	httpStatus int
	response   string
}

type requestResult struct {
	Result string `json:"result"`
}

//TestController test endpoints
func TestController(t *testing.T) {
	server := server.New(server.Config{})

	group := server.NewGroup("")
	fizzbuzzService := NewMockService()
	statsService := statistics.NewStatsMockService()
	RegisterHandlers(group, fizzbuzzService, statsService)

	tests := []testCasesController{
		{"Bad int1", postRequest{Int1: -1, Int2: 4, Limit: 10, Str1: "a", Str2: "b"}, 400, ""},
		{"Bad int2", postRequest{Int1: 2, Int2: -1, Limit: 10, Str1: "a", Str2: "b"}, 400, ""},
		{"Bad limit", postRequest{Int1: 2, Int2: 4, Limit: 1000000, Str1: "a", Str2: "b"}, 400, ""},
		{"Bad str1", postRequest{Int1: 2, Int2: 4, Limit: 10, Str1: "", Str2: "b"}, 400, ""},
		{"Bad str2", postRequest{Int1: 2, Int2: 4, Limit: 10, Str1: "a", Str2: ""}, 400, ""},
		{"Limit too low", postRequest{Int1: 2, Int2: 4, Limit: 1, Str1: "a", Str2: "b"}, 400, ""},
		{"Returns ok", postRequest{Int1: 2, Int2: 4, Limit: 10, Str1: "a", Str2: "b"}, 200, "mock"},
		{"Returns ok", postRequest{Int1: 3, Int2: 5, Limit: 16, Str1: "fizz", Str2: "buzz"}, 200, "mock"},
	}

	for _, tc := range tests {
		response := requestResult{}
		requestByte, _ := json.Marshal(tc.request)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/fizzbuzz", bytes.NewReader(requestByte))
		server.Router.ServeHTTP(w, req)

		logger.Info(w.Result())
		assert.Equal(t, tc.httpStatus, w.Code)

		err := json.Unmarshal([]byte(w.Body.String()), &response)
		assert.NoError(t, err)

		assert.Equal(t, tc.response, response.Result)
	}
}
