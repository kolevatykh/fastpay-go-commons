package requests

type ConfirmInvitationRequest struct {
	Id      string  `json:"id" validate:"required,uuid"`
	Address string  `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Amount  int64   `json:"amount" validate:"range(0|9223372036854775807)"`
	MsgHash string  `json:"msgHash" validate:"required"`
	Sig     SignDto `json:"sig" validate:"required"`
	Exp     int64   `json:"exp" valid:"required~ErrorTimestampNotPassed"`
}
