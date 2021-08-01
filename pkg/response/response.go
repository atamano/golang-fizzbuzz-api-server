package response

import (
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
)

type ErrorResponse struct {
	Error       string   `json:"-"`
	Message     string   `json:"message"`
	Validations []string `json:"validations"`
}

func BuildErrorReponse(err error, message string) ErrorResponse {
	validationErrors := []string{}

	if _, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprint(fieldErr))
		}
	}

	return ErrorResponse{
		Error:       err.Error(),
		Message:     message,
		Validations: validationErrors,
	}
}

type PaginatedResponse struct {
	Results interface{} `json:"results"`
	Page    int         `json:"page"`
	Limit   int         `json:"limit"`
}

func BuildPaginatedResponse(res interface{}, page int, limit int) PaginatedResponse {
	return PaginatedResponse{
		Results: res,
		Page:    page,
		Limit:   limit,
	}
}
