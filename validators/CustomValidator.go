package validators

import "github.com/go-playground/validator/v10"

type NewCustomValidator struct {
	Validator *validator.Validate
}

func (cv *NewCustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
