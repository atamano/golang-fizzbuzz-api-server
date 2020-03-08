package statistics

import (
	"github.com/go-pg/pg"
)

//Repository for fizzbuzz
type Repository interface {
	Increment(key string) (*FizzbuzzRequestsStats, error)
	GetMostUsedRequest() (*FizzbuzzRequestsStats, error)
	Create(key string, params []byte) (*FizzbuzzRequestsStats, error)
	Get(key string) (*FizzbuzzRequestsStats, error)
}

type repository struct {
	db *pg.DB
}

//NewRepository initialization
func NewRepository(db *pg.DB) Repository {
	return repository{db}
}

func (r repository) Get(key string) (*FizzbuzzRequestsStats, error) {
	requestStats := new(FizzbuzzRequestsStats)
	err := r.db.Model(requestStats).Where("key = ?", key).Select()

	return requestStats, err
}

func (r repository) Increment(key string) (*FizzbuzzRequestsStats, error) {
	requestStats := new(FizzbuzzRequestsStats)
	_, err := r.db.Model(requestStats).Set("counter = counter + 1").Where("key = ?", key).Update()

	return requestStats, err
}

func (r repository) Create(key string, params []byte) (*FizzbuzzRequestsStats, error) {
	requestStats := FizzbuzzRequestsStats{Key: key, Params: []byte(params), Counter: 1}
	_, err := r.db.Model(&requestStats).Returning("*").Insert()

	return &requestStats, err
}

func (r repository) GetMostUsedRequest() (*FizzbuzzRequestsStats, error) {
	requestStats := new(FizzbuzzRequestsStats)
	err := r.db.Model(requestStats).Order("counter DESC").First()
	return requestStats, err
}
