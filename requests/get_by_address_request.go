package requests

type GetByAddressRequest struct {
	Address string `json:"address" validate:"required,validHex40"`
}
