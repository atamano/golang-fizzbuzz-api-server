package statistics

import (
	"github.com/atamano/fizz-buzz/pkg/request"
)

//Service for fizzbuzz
type Service interface {
	IncrementRequestCount(request request.Request) (fizzbuzzRequestsStats, error)
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

func (s service) IncrementRequestCount(request request.Request) (fizzbuzzRequestsStats, error) {
	key := request.ToStr()
	params := request.ToJSON()
	result, err := s.repository.get(key)
	if err == nil {
		result, err = s.repository.increment(key)
	} else {
		result, err = s.repository.create(key, params)
	}
	return result, err
}

func (s service) GetMostUsedRequest() (fizzbuzzRequestsStats, error) {
	return s.repository.getMostUsedRequest()
}
