package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"

type UpdateStateClientBankRequest struct {
	State state_enum.State `json:"state" validate:"required"`
}
