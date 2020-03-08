package fizzbuzz

import (
	"strconv"
	"strings"
)

//Service for fizzbuzz
type Service interface {
	Compute(args *PostRequest) string
}

type service struct{}

//NewService initialization
func NewService() Service {
	return service{}
}

//Compute fizzbuzz
func (s service) Compute(params *PostRequest) string {

	result := make([]string, params.Limit)

	for i := 1; i <= params.Limit; i++ {
		if i%params.Int1 == 0 && i%params.Int2 == 0 {
			result[i-1] = params.Str1 + params.Str2
		} else if i%params.Int1 == 0 {
			result[i-1] = params.Str1
		} else if i%params.Int2 == 0 {
			result[i-1] = params.Str2
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}

	return strings.Join(result, ",")
}
