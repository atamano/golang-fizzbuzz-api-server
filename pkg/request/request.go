package request

// Request interface for requests
type Request interface {
	Validate() error
	ToStr() string
	ToJSON() []byte
}
