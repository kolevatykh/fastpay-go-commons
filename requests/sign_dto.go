package requests

type SignDto struct {
	R string `json:"r" validate:"required,validHex64"`
	S string `json:"s" validate:"required,validHex64"`
	V int    `json:"v" validate:"required,gte=27,lte=28"`
}
