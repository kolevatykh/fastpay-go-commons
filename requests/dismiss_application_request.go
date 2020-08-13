package requests

type DismissApplicationRequest struct {
	Id               string  `json:"id" validate:"required,uuid"`
	AddressInitiator string  `json:"addressInitiator" validate:"required,validHex40"`
	Exp              int64   `json:"exp" validate:"required"`
	MsgHash          string  `json:"msgHash" validate:"required"`
	Sig              SignDto `json:"sig" validate:"required,dive"`
}
