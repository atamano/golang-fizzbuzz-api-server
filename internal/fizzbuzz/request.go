package fizzbuzz

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/sirupsen/logrus"
)

//PostRequest Post request parameters
type PostRequest struct {
	Int1  int    `json:"int1" binding:"required"`
	Int2  int    `json:"int2" binding:"required"`
	Limit int    `json:"limit" binding:"required"`
	Str1  string `json:"str1" binding:"required"`
	Str2  string `json:"str2" binding:"required"`
}

//Validate request
func (a PostRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Int1, validation.Required, validation.Min(1), validation.Max(a.Limit)),
		validation.Field(&a.Int2, validation.Required, validation.Min(1), validation.Max(a.Limit)),
		validation.Field(&a.Limit, validation.Required, validation.Min(1), validation.Max(100000)),
		validation.Field(&a.Str1, validation.Required, validation.Length(1, 100)),
		validation.Field(&a.Str2, validation.Required, validation.Length(1, 100)),
	)
}

//GetRequestKey for unique key
func (a PostRequest) GetRequestKey() string {

	// to avoid injection and separator character in params
	str1 := base64.StdEncoding.EncodeToString([]byte(a.Str1))
	str2 := base64.StdEncoding.EncodeToString([]byte(a.Str2))

	return fmt.Sprintf("fizzbuzz|%d|%d|%d|%s|%s", a.Int1, a.Int2, a.Limit, str1, str2)
}

//ToJSON struct to json
func (a PostRequest) ToJSON() []byte {

	b, err := json.Marshal(a)

	if err != nil {
		logrus.WithError(err).Fatal("Cannot encode to json")
	}

	return b
}
