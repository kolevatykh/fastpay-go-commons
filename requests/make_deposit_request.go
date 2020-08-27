package requests

type MakeDepositRequest struct {
	Id          string  `json:"id" valid:"required,uuid"`
	AddressFrom string  `json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	AddressTo   string  `json:"addressTo" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Amount      int64   `json:"amount" valid:"required~ErrorAmountNotPassed,range(0|9223372036854775807)~ErrorAmountNegative"`
	MsgHash     string  `json:"msgHash" valid:"required"`
	Sig         SignDto `json:"sig" valid:"required"`
	Exp         int64   `json:"exp" valid:"required~ErrorTimestampNotPassed"`
}
