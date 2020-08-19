package requests

type DismissApplicationRequest struct {
	Id               string  `json:"id" valid:"required~60304,uuid"`
	AddressInitiator string  `json:"addressInitiator" valid:"required~60302,validHex40~60301"`
	Exp              int64   `json:"exp" valid:"required~60332"`
	MsgHash          string  `json:"msgHash" valid:"required"`
	Sig              SignDto `json:"sig" valid:"required"`
}
