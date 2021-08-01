package statistics

import (
	"encoding/json"
	"fmt"
)

type Service interface {
	IncrementRequestCount(params interface{}) (fizzbuzzRequestsStats, error)
	GetMostUsedRequest() (fizzbuzzRequestsStats, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s service) IncrementRequestCount(params interface{}) (fizzbuzzRequestsStats, error) {
	key := fmt.Sprintf("%v", params)

	b, err := json.Marshal(params)
	if err != nil {
		return fizzbuzzRequestsStats{}, err
	}

	result, err := s.repository.Get(key)
	if err == nil {
		result, err = s.repository.Increment(key)
	} else {
		result, err = s.repository.Create(key, b)
	}

	return result, err
}

func (s service) GetMostUsedRequest() (fizzbuzzRequestsStats, error) {
	return s.repository.GetMostUsedRequest()
}
