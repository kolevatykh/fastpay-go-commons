package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"

type CreateBankRequest struct {
	Address string           `json:"address"  validate:"required,validHex40"`
	State   state_enum.State `json:"state"`
	Name    string           `json:"name"`
	BIK     string           `json:"bik"`
	MspId   string           `json:"mspId"`
	IsOwner bool             `json:"isOwner"`
	Conf    string           `json:"conf"`
}

func (createBank *CreateBankRequest) SetDefaults() {
	if createBank.State == state_enum.Undefined {
		createBank.State = state_enum.Available
	}
}
