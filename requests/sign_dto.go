package requests

type SignDto struct {
	R string `json:"r" valid:"validHex64~60323"`
	S string `json:"s" valid:"validHex64~60325"`
	V int    `json:"v" valid:"range(27|28)~60327"`
}
