package statistics

import "github.com/atamano/fizz-buzz/pkg/request"

type statsMockService struct{}

func (m statsMockService) IncrementRequestCount(request request.Request) (fizzbuzzRequestsStats, error) {
	return fizzbuzzRequestsStats{}, nil
}

func (m statsMockService) GetMostUsedRequest() (fizzbuzzRequestsStats, error) {
	return fizzbuzzRequestsStats{}, nil
}

//NewStatsMockService for testing stats
func NewStatsMockService() Service {
	return statsMockService{}
}
