package statistics

import (
	"github.com/go-pg/pg/v10/orm"
)

//Repository for fizzbuzz
type Repository interface {
	increment(key string) (fizzbuzzRequestsStats, error)
	getMostUsedRequest() (fizzbuzzRequestsStats, error)
	create(key string, params []byte) (fizzbuzzRequestsStats, error)
	get(key string) (fizzbuzzRequestsStats, error)
}

type repository struct {
	db orm.DB
}

//NewRepository initialization
func NewRepository(db orm.DB) Repository {
	return repository{db}
}

func (r repository) get(key string) (fizzbuzzRequestsStats, error) {
	requestStats := fizzbuzzRequestsStats{}
	err := r.db.Model(&requestStats).Where("key = ?", key).Select()

	return requestStats, err
}

func (r repository) increment(key string) (fizzbuzzRequestsStats, error) {
	requestStats := fizzbuzzRequestsStats{}
	_, err := r.db.Model(&requestStats).Set("counter = counter + 1").Where("key = ?", key).Update()

	return requestStats, err
}

func (r repository) create(key string, params []byte) (fizzbuzzRequestsStats, error) {
	requestStats := fizzbuzzRequestsStats{Key: key, Params: []byte(params), Counter: 1}
	_, err := r.db.Model(&requestStats).Returning("*").Insert()

	return requestStats, err
}

func (r repository) getMostUsedRequest() (fizzbuzzRequestsStats, error) {
	requestStats := fizzbuzzRequestsStats{}
	err := r.db.Model(&requestStats).Order("counter DESC").First()
	return requestStats, err
}
