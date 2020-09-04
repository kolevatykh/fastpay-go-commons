package requests

type ConfirmInvitationRequest struct {
	Id               string  `json:"id" valid:"required,uuid"`
	Address          string  `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Amount           int64   `json:"amount" valid:"range(0|9223372036854775807)"`
	ObligatoryAmount int64   `json:"obligatoryAmount" valid:"range(0|9223372036854775807)"`
	MsgHash          string  `json:"msgHash" valid:"required"`
	Sig              SignDto `json:"sig" valid:"required"`
	Exp              int64   `json:"exp" valid:"required~ErrorTimestampNotPassed"`
}
