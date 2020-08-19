package validation

import (
	"github.com/asaskevich/govalidator"
	"math"
	"regexp"
)

var (
	hex64          = "^[a-f0-9]{64}$"
	hex64Regex     = regexp.MustCompile(hex64)
	hex40          = "^[a-f0-9]{40}$"
	hex40Regex     = regexp.MustCompile(hex40)
	hex40or64      = "^([a-f0-9]{40}|[a-f0-9]{64})$"
	hex40or64Regex = regexp.MustCompile(hex40or64)
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

	govalidator.ParamTagMap["min"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		if len(params) == 1 {
			return govalidator.InRange(str, params[0], math.MaxInt64)
		}
		return false
	})

	govalidator.ParamTagMap["max"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		if len(params) == 1 {
			return govalidator.InRange(str, math.MinInt64, params[0])
		}
		return false
	})

	govalidator.ParamTagMap["gte"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		if len(params) == 1 {
			return govalidator.InRange(str, params[0], math.MaxInt64)
		}
		return false
	})

	govalidator.ParamTagMap["lte"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		if len(params) == 1 {
			return govalidator.InRange(str, math.MinInt64, params[0])
		}
		return false
	})

}

func ValidateStruct(s interface{}) (bool, error) {
	return govalidator.ValidateStruct(s)
}
