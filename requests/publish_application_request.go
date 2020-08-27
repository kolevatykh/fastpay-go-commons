package requests

type PublishApplicationRequest struct {
	Id           string  `json:"id" valid:"required,uuid"`
	AddressOwner string  `json:"addressOwner" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	MsgHash      string  `json:"msgHash" valid:"required"`
	Sig          SignDto `json:"sig" valid:"required"`
	Exp          int64   `json:"exp" valid:"required~ErrorTimestampNotPassed"`
}
