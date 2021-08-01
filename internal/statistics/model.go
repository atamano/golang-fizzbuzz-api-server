package statistics

import "encoding/json"

type fizzbuzzRequestsStats struct {
	ID      int             `json:"id"`
	Key     string          `json:"key"`
	Params  json.RawMessage `json:"params"`
	Counter int             `json:"counter"`
}
