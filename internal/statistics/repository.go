package statistics

import (
	"github.com/go-pg/pg/v10/orm"
)

type Repository interface {
	Increment(key string) (fizzbuzzRequestsStats, error)
	GetMostUsedRequest() (fizzbuzzRequestsStats, error)
	Create(key string, params []byte) (fizzbuzzRequestsStats, error)
	Get(key string) (fizzbuzzRequestsStats, error)
}

type repository struct {
	db orm.DB
}

func NewRepository(db orm.DB) Repository {
	return repository{db}
}

func (r repository) Get(key string) (fizzbuzzRequestsStats, error) {
	requestStats := fizzbuzzRequestsStats{}
	err := r.db.Model(&requestStats).Where("key = ?", key).Select()

	return requestStats, err
}

func (r repository) Increment(key string) (fizzbuzzRequestsStats, error) {
	requestStats := fizzbuzzRequestsStats{}
	_, err := r.db.Model(&requestStats).Set("counter = counter + 1").Where("key = ?", key).Update()

	return requestStats, err
}

func (r repository) Create(key string, params []byte) (fizzbuzzRequestsStats, error) {
	requestStats := fizzbuzzRequestsStats{Key: key, Params: []byte(params), Counter: 1}
	_, err := r.db.Model(&requestStats).Returning("*").Insert()

	return requestStats, err
}

func (r repository) GetMostUsedRequest() (fizzbuzzRequestsStats, error) {
	requestStats := fizzbuzzRequestsStats{}
	err := r.db.Model(&requestStats).Order("counter DESC").First()
	return requestStats, err
}
