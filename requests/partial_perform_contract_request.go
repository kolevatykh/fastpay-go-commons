package requests

type PartialPerformContractRequest struct {
	Id                    string  `json:"id" valid:"required,uuid"`
	Address               string  `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	MsgHash               string  `json:"msgHash" valid:"required"`
	Sig                   SignDto `json:"sig" valid:"required"`
	Exp                   int64   `json:"exp" valid:"required~ErrorTimestampNotPassed"`
	ActualAmountInitiator int64   `json:"actualAmountInitiator" valid:"required,range(0|9223372036854775807)~ErrorAmountNegative"`
	ActualAmountAcceptor  int64   `json:"actualAmountAcceptor" valid:"required,range(0|9223372036854775807)~ErrorAmountNegative"`
}
