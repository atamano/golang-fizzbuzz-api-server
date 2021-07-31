package statistics

//Service for fizzbuzz
type Service interface {
	IncrementRequestCount(rkey string, params []byte) (fizzbuzzRequestsStats, error)
	GetMostUsedRequest() (fizzbuzzRequestsStats, error)
}

type service struct {
	repository Repository
}

//NewService initialization
func NewService(repository Repository) Service {
	return &service{repository}
}

//StatsRequest interface

func (s service) IncrementRequestCount(key string, params []byte) (fizzbuzzRequestsStats, error) {
	result, err := s.repository.Get(key)
	if err == nil {
		result, err = s.repository.Increment(key)
	} else {
		result, err = s.repository.Create(key, params)
	}
	return result, err
}

func (s service) GetMostUsedRequest() (fizzbuzzRequestsStats, error) {
	return s.repository.GetMostUsedRequest()
}
