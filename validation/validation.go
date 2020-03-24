package validation

import (
	"regexp"

	"github.com/go-playground/validator"
)

var (
	Validate       = validator.New()
	hex64          = "^[a-f0-9]{64}$"
	hex64Regex     = regexp.MustCompile(hex64)
	hex40          = "^[a-f0-9]{40}$"
	hex40Regex     = regexp.MustCompile(hex40)
	hex40or64      = "^([a-f0-9]{40}|[a-f0-9]{64})$"
	hex40or64Regex = regexp.MustCompile(hex40or64)
)

func init() {
	Validate.RegisterValidation("validHex64", func(fl validator.FieldLevel) bool {
		return hex64Regex.MatchString(fl.Field().String())
	})

	Validate.RegisterValidation("validHex40", func(fl validator.FieldLevel) bool {
		return hex40Regex.MatchString(fl.Field().String())
	})

	Validate.RegisterValidation("validHex40or64", func(fl validator.FieldLevel) bool {
		return hex40or64Regex.MatchString(fl.Field().String())
	})
}
