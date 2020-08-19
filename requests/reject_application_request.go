package requests

type RejectApplicationRequest struct {
	Id      string  `json:"id" valid:"required,uuid"`
	Address string  `json:"address" valid:"required,validHex40"`
	MsgHash string  `json:"msgHash" valid:"required"`
	Sig     SignDto `json:"sig" valid:"required"`
	Exp     int64   `json:"exp" valid:"required"`
}
