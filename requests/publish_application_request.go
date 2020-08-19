package requests

type PublishApplicationRequest struct {
	Id               string  `json:"id" valid:"required,uuid"`
	AddressInitiator string  `json:"addressInitiator" valid:"required~60302,validHex40~60301"`
	Exp              int64   `json:"exp" valid:"required~60332"`
	MsgHash          string  `json:"msgHash" valid:"required"`
	Sig              SignDto `json:"sig" valid:"required"`
}
