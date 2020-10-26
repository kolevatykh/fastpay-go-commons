package requests

type AddRegulatorRequest struct {
	TechnicalSignRequest
	Address string `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}
