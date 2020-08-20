package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"

type CreateBankRequest struct {
	Address string           `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	State   state_enum.State `json:"state"`
	Name    string           `json:"name" valid:"required"`
	BIK     string           `json:"bik" valid:"required"`
	MspId   string           `json:"mspId" valid:"required"`
	IsOwner bool             `json:"isOwner"`
	Conf    string           `json:"conf"`
}

func (createBank *CreateBankRequest) SetDefaults() {
	if createBank.State == state_enum.Undefined {
		createBank.State = state_enum.Available
	}
}
