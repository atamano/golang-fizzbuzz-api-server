package statistics

type statsMockService struct{}

func (m statsMockService) IncrementRequestCount(params interface{}) (fizzbuzzRequestsStats, error) {
	return fizzbuzzRequestsStats{}, nil
}

func (m statsMockService) GetMostUsedRequest() (fizzbuzzRequestsStats, error) {
	return fizzbuzzRequestsStats{}, nil
}

//NewStatsMockService for testing stats
func NewStatsMockService() Service {
	return statsMockService{}
}
