package requests

type SignDto struct {
	R string `json:"r" validate:"omitempty,validHex64"`
	S string `json:"s" validate:"omitempty,validHex64"`
	V int    `json:"v" validate:"omitempty,gte=27,lte=28"`
}
