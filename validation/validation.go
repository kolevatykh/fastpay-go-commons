package validation

import (
	"fmt"
	"github.com/go-playground/validator"
	"regexp"
)

var (
	Validate   = validator.New()
	hex64      = "^[a-f0-9]{64}$" // regex that compiles
	hex64Regex = regexp.MustCompile(hex64)
	hex40      = "^[a-f0-9]{40}$" // regex that compiles
	hex40Regex = regexp.MustCompile(hex40)
)

func init() {
	fmt.Println("|||--init validator--||")
	Validate.RegisterValidation("validHex64", func(fl validator.FieldLevel) bool {
		return hex64Regex.MatchString(fl.Field().String())
	})

	Validate.RegisterValidation("validHex40", func(fl validator.FieldLevel) bool {
		return hex40Regex.MatchString(fl.Field().String())
	})
}
