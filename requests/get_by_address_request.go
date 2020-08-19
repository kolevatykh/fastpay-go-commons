package requests

type GetByAddressRequest struct {
	Address string `json:"address" valid:"required~60302,validHex40~60301"`
}
