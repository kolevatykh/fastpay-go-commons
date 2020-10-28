package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"

type UpdateAccountStateRequest struct {
	Address string           `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	State   state_enum.State `json:"state" valid:"required"`
}
