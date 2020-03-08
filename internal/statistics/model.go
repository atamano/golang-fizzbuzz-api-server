package statistics

import "encoding/json"

//FizzbuzzRequestsStats counts requests
type FizzbuzzRequestsStats struct {
	tableName struct{}        `pg:"fizzbuzz_requests_stats"`
	ID        int             `pg:"id" json:"id"`
	Key       string          `pg:"key,unique" json:"key"`
	Params    json.RawMessage `pg:"params" json:"params"`
	Counter   int             `pg:"counter,default:0" json:"counter"`
}
