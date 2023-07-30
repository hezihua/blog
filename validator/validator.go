package validator

import "github.com/go-playground/validator/v10"

var (
	validate *validator.Validate
)


func V() *validator.Validate {
	return validate
}

func init() {
	validate = validator.New()


}