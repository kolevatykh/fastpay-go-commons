package requests

type PublishApplicationRequest struct {
	Id               string  `json:"id" valid:"required,uuid"`
	AddressInitiator string  `json:"addressInitiator" valid:"required,validHex40"`
	Exp              int64   `json:"exp" valid:"required"`
	MsgHash          string  `json:"msgHash" valid:"required"`
	Sig              SignDto `json:"sig" valid:"required"`
}
