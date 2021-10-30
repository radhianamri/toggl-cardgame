package validator

import (
	"github.com/go-playground/validator/v10"
)

var (
	valid *validator.Validate
)

func Init() {
	valid = validator.New()
}

func ValidateInput(data interface{}, tags string) error {
	return valid.Var(data, tags)
}

func ValidateStructInput(v interface{}) error {
	return valid.Struct(v)
}
