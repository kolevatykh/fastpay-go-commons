package requests

type TechnicalSignRequest struct {
	TechnicalMsgHash string  `json:"technicalMsgHash" valid:"required"`
	TechnicalAddress string  `json:"technicalAddress" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	TechnicalSig     SignDto `json:"technicalSig" valid:"required"`
}
