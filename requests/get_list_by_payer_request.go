package requests

type GetListByPayerRequest struct {
	Payer string `json:"payer" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}
