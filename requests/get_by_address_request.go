package requests

type GetByAddressRequest struct {
	Address string `json:"address" valid:"required,validHex40"`
}
