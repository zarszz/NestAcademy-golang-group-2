package params

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

func Validate(u interface{}) error {
	err := validator.New().Struct(u)
	if err == nil {
		return nil
	}
	myErr := err.(validator.ValidationErrors)
	errString := ""
	for _, e := range myErr {
		errString += e.Field() + " is " + e.Tag()
	}
	return errors.New(errString)
}
