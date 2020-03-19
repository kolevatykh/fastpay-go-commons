package requests

type AddIdentifierRequest struct {
	Address    string  `json:"address" validate:"required,validHex40"`
	Identifier string  `json:"identifier" validate:"required,validHex64"`
	MsgHash    string  `json:"msgHash" validate:"required"`
	Sig        SignDto `json:"sig" validate:"required,dive"`
}
