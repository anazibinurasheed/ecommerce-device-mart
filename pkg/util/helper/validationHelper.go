package helper

import "github.com/go-playground/validator/v10"

func ValidateInputStruct(s interface{}) error {
	validate := validator.New()

	return validate.Struct(s)
}
