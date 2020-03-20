package requests

type AddRegulatorRequest struct {
	Address string `json:"address" validate:"required,validHex40"`
}
