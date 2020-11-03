package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"

type UpdateStateClientBankRequest struct {
	TechnicalSignRequest
	State   state_enum.State `json:"state" valid:"required"`
	Address string           `json:"address" valid:"required~ErrorBankAddressNotPassed"`
}
