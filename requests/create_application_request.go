package requests

type CreateApplicationRequest struct {
	Id                       string  `json:"id" valid:"required,uuid"`
	AddressInitiator         string  `json:"addressInitiator" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	AddressAcceptor          string  `json:"addressAcceptor" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	AmountVotesCloseContract int     `json:"amountVotesCloseContract" valid:"required,range(0|9223372036854775807)"`
	MsgHash                  string  `json:"msgHash" valid:"required"`
	Sig                      SignDto `json:"sig" valid:"required"`
	Exp                      int64   `json:"exp" valid:"required~ErrorTimestampNotPassed""`
}
