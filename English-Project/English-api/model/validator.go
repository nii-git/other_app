package model

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()
	v.RegisterValidation("date_validation", Date_validation)
	v.RegisterValidation("limit_validation", Limit_validation)
	v.RegisterValidation("page_validation", Page_validation)
	return &Validator{Validator: *v}
}

func Date_validation(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	return regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$`).Match([]byte(date))
}

func Limit_validation(fl validator.FieldLevel) bool {
	limit := fl.Field().Int()
	// todo: MAX_LIMITは環境変数に入れたい
	return (limit >= 0) && (limit < 1000)
}

func Page_validation(fl validator.FieldLevel) bool {
	page := fl.Field().Int()
	return (page >= 0)
}
