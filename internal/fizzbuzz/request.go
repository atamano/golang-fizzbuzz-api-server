package fizzbuzz

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/atamano/fizz-buzz/pkg/logger"
)

//postRequest Post request parameters
type postRequest struct {
	Int1  int    `json:"int1" binding:"required,gte=1,lt=10000"`
	Int2  int    `json:"int2" binding:"required,gte=1,lt=10000"`
	Limit int    `json:"limit" binding:"required,gte=1,lt=10000"`
	Str1  string `json:"str1" binding:"required,min=1,max=100"`
	Str2  string `json:"str2" binding:"required,min=1,max=100"`
}

//ToStr for unique key
func (a postRequest) ToStr() string {

	// to avoid injection and separator character in params
	str1 := base64.StdEncoding.EncodeToString([]byte(a.Str1))
	str2 := base64.StdEncoding.EncodeToString([]byte(a.Str2))

	return fmt.Sprintf("fizzbuzz|%d|%d|%d|%s|%s", a.Int1, a.Int2, a.Limit, str1, str2)
}

//ToBytes converts struct to bytes
func (a postRequest) ToBytes() []byte {

	b, err := json.Marshal(a)

	if err != nil {
		logger.Fatal("Cannot encode to json", err.Error())
	}

	return b
}
