package requests

type AcceptApplicationRequest struct {
	Id              string  `json:"id" valid:"required,uuid"`
	AddressAcceptor string  `json:"addressAcceptor" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	MsgHash         string  `json:"msgHash" valid:"required"`
	Sig             SignDto `json:"sig" valid:"required"`
	Exp             int64   `json:"exp" valid:"required~ErrorTimestampNotPassed"`
}
