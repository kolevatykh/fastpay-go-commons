package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"

type UpdateStateRequest struct {
	Address string           `json:"address" validate:"required, validHex40"`
	State   state_enum.State `json:"state" validate:"required"`
}
