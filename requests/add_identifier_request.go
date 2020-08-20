package requests

type AddIdentifierRequest struct {
	Address    string  `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Identifier string  `json:"identifier" valid:"required~ErrorIdentifierNotPassed,validHex64"`
	MsgHash    string  `json:"msgHash" valid:"required"`
	Sig        SignDto `json:"sig" valid:"required"`
}
