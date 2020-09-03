package requests

type RemoveArbitratorRequest struct {
	Address string `json:"address" valid:"validHex40or64~ErrorAddressNotFollowingRegex"`
}
