package fizzbuzz

import (
	"testing"

	"github.com/atamano/fizz-buzz/internal/statistics"
	"github.com/stretchr/testify/assert"
)

type statsMockService struct{}

func (m statsMockService) IncrementRequestCount(request statistics.StatsRequest) (*statistics.FizzbuzzRequestsStats, error) {
	return nil, nil
}

func (m statsMockService) GetMostUsedRequest() (*statistics.FizzbuzzRequestsStats, error) {
	return nil, nil
}

//newStatsMockService for testing stats
func newStatsMockService() statistics.Service {
	return statsMockService{}
}

type mockService struct{}

func (s mockService) Compute(params *PostRequest) string {
	return "mock"
}

//newMockService for testing fizzbuz
func newMockService() Service {
	return mockService{}
}

type testCasesService struct {
	request  PostRequest
	response string
}

//TestService test endpoints
func TestService(t *testing.T) {

	service := NewService()

	tests := []testCasesService{
		{PostRequest{Int1: 2, Int2: 4, Limit: 10, Str1: "a", Str2: "b"}, "1,a,3,ab,5,a,7,ab,9,a"},
		{PostRequest{Int1: 3, Int2: 5, Limit: 16, Str1: "fizz", Str2: "buzz"}, "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16"},
		{PostRequest{Int1: 5, Int2: 5, Limit: 15, Str1: "a", Str2: "b"}, "1,2,3,4,ab,6,7,8,9,ab,11,12,13,14,ab"},
		{PostRequest{Int1: 1, Int2: 1, Limit: 15, Str1: "a", Str2: "b"}, "ab,ab,ab,ab,ab,ab,ab,ab,ab,ab,ab,ab,ab,ab,ab"},
	}

	for _, tc := range tests {
		result := service.Compute(&tc.request)
		assert.Equal(t, tc.response, result)
	}
}
