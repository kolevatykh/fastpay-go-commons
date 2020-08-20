package requests

type AddRegulatorRequest struct {
	Address string `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}
