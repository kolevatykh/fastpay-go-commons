package requests

type DismissApplicationRequest struct {
	Id               string  `json:"id" valid:"required,uuid"`
	AddressInitiator string  `json:"addressInitiator" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Exp              int64   `json:"exp" valid:"required~ErrorTimestampNotPassed"`
	MsgHash          string  `json:"msgHash" valid:"required"`
	Sig              SignDto `json:"sig" valid:"required"`
}
