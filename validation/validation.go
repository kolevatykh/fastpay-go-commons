package validation

import (
	"github.com/asaskevich/govalidator"
	"regexp"
)

var (
	hex64          = "^[a-f0-9]{64}$"
	hex64Regex     = regexp.MustCompile(hex64)
	hex40          = "^[a-f0-9]{40}$"
	hex40Regex     = regexp.MustCompile(hex40)
	hex40or64      = "^([a-f0-9]{40}|[a-f0-9]{64})$"
	hex40or64Regex = regexp.MustCompile(hex40or64)
	hex62or64      = "^([a-f0-9]{62}|[a-f0-9]{64})$"
	hex62or64Regex = regexp.MustCompile(hex62or64)
)

func init() {
	govalidator.TagMap["validHex64"] = govalidator.Validator(func(str string) bool {
		return hex64Regex.MatchString(str)
	})

	govalidator.TagMap["validHex40"] = govalidator.Validator(func(str string) bool {
		return hex40Regex.MatchString(str)
	})

	govalidator.TagMap["validHex40or64"] = govalidator.Validator(func(str string) bool {
		return hex40or64Regex.MatchString(str)
	})

	govalidator.TagMap["validHex62or64"] = govalidator.Validator(func(str string) bool {
		return hex62or64Regex.MatchString(str)
	})
}

func ValidateStruct(s interface{}) (bool, error) {
	return govalidator.ValidateStruct(s)
}
