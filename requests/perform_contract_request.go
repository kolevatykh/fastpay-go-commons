package requests

type PerformContractRequest struct {
	Id      string  `json:"id" valid:"required,uuid"`
	Address string  `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	MsgHash string  `json:"msgHash" valid:"required"`
	Sig     SignDto `json:"sig" valid:"required"`
	Exp     int64   `json:"exp" valid:"required~ErrorTimestampNotPassed"`
}
