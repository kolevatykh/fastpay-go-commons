package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"

type UpdateStateRequest struct {
	Address string           `json:"address" valid:"required~60302,validHex40~60301"`
	State   state_enum.State `json:"state" valid:"required"`
}
