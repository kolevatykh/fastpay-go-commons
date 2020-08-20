package requests

type SignDto struct {
	R string `json:"r" valid:"validHex64~ErrorSigRNotFolowingRegex"`
	S string `json:"s" valid:"validHex64~ErrorSigSNotFolowingRegex"`
	V int    `json:"v" valid:"range(27|28)~ErrorSigVIncorrect"`
}
