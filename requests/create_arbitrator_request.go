package requests

type CreateArbitratorRequest struct {
	Address string `json:"address" valid:"validHex40or64~ErrorAddressNotFollowingRegex"`
	Name    string `json:"name" valid:"required"`
}
