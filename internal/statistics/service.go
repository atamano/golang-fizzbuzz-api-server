package statistics

import "github.com/sirupsen/logrus"

//Service for fizzbuzz
type Service interface {
	IncrementRequestCount(request StatsRequest) (*FizzbuzzRequestsStats, error)
	GetMostUsedRequest() (*FizzbuzzRequestsStats, error)
}

type service struct {
	repository Repository
}

//NewService initialization
func NewService(repository Repository) Service {
	return &service{repository}
}

//StatsRequest interface
type StatsRequest interface {
	GetRequestKey() string
	ToJSON() []byte
}

func (s service) IncrementRequestCount(request StatsRequest) (*FizzbuzzRequestsStats, error) {
	key := request.GetRequestKey()
	params := request.ToJSON()
	result, err := s.repository.Get(key)
	if err == nil {
		result, err = s.repository.Increment(key)
	} else {
		result, err = s.repository.Create(key, params)
	}
	if err != nil {
		logrus.WithError(err)
	}
	return result, err
}

func (s service) GetMostUsedRequest() (*FizzbuzzRequestsStats, error) {
	return s.repository.GetMostUsedRequest()
}
