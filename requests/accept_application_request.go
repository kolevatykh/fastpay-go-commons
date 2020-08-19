package requests

type AcceptApplicationRequest struct {
	Id              string  `json:"id" valid:"required~60304,uuid"`
	AddressAcceptor string  `json:"addressAcceptor" valid:"required~60302,validHex40~60301"`
	MsgHash         string  `json:"msgHash" valid:"required"`
	Sig             SignDto `json:"sig" valid:"required"`
	Exp             int64   `json:"exp" valid:"required~60332"`
}
