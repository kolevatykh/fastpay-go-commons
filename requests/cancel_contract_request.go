package requests

type CancelContractRequest struct {
	Id      string  `json:"id" valid:"required~60304,uuid"`
	Address string  `json:"address" valid:"required~60302,validHex40~60303"`
	MsgHash string  `json:"msgHash" valid:"required"`
	Sig     SignDto `json:"sig" valid:"required"`
	Exp     int64   `json:"exp" valid:"required~60332"`
}
