package requests

type RejectApplicationRequest struct {
	Id              string  `json:"id" validate:"required,uuid"`
	AddressAcceptor string  `json:"addressAcceptor" validate:"required,validHex40"`
	MsgHash         string  `json:"msgHash" validate:"required"`
	Sig             SignDto `json:"sig" validate:"required,dive"`
	Exp             int64   `json:"exp" validate:"required"`
}
