package requests

type GetPageByPayerRequest struct {
	GetPageRequest
	Payer string `json:"payer" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}
