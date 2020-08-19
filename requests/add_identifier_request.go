package requests

type AddIdentifierRequest struct {
	Address    string  `json:"address" valid:"required~60302,validHex40~60303"`
	Identifier string  `json:"identifier" valid:"required~60304,validHex64"`
	MsgHash    string  `json:"msgHash" valid:"required"`
	Sig        SignDto `json:"sig" valid:"required"`
}
