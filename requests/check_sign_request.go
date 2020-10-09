package requests

type CheckSignRequest struct {
	MsgHash string  `json:"msgHash" valid:"required"`
	Exp     int64   `json:"exp" valid:"required~ErrorTimestampNotPassed"`
	Sig     SignDto `json:"sig" valid:"required"`
}
