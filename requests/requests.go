package requests

import (
	account_type "github.com/SolarLabRU/fastpay-go-commons/enums/account-type"
	identity_type "github.com/SolarLabRU/fastpay-go-commons/enums/identity-type"
	juridical_type "github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state"
)

type CreateAccountRequest struct {
	Address       string                       `json:"address" validate:"required,validHex40"`
	State         state.State                  `json:"state"`
	CurrencyCode  int                          `json:"currencyCode" validate:"required,gte=0,lte=999"`
	JuridicalType juridical_type.JuridicalType `json:"juridicalType"`
	IdentityType  identity_type.IdentityType   `json:"identityType"`
	Type          account_type.AccountType     `json:"type"`
	Identifiers   []string                     `json:"identifiers" validate:"required,dive,validHex64"`
	Timestamp     int64                        `json:"timestamp" validate:"required"`
	PublicKey     string                       `json:"publicKey"`
	MsgHash       string                       `json:"msgHash" validate:"required"`
	Sig           SignDto                      `json:"sig" validate:"required,dive"`
}

func (createAccount *CreateAccountRequest) SetDefaults() {
	if createAccount.IdentityType == identity_type.Undefined {
		createAccount.IdentityType = identity_type.None
	}
	if createAccount.State == state.Undefined {
		createAccount.State = state.Available
	}
	if createAccount.Type == account_type.Undefined {
		createAccount.Type = account_type.Client
	}
}

type CreateBankRequest struct {
	Address string      `json:"address"  validate:"required,validHex40"`
	State   state.State `json:"state"`
	Name    string      `json:"name"`
	BIK     string      `json:"bik"`
	MspId   string      `json:"mspId"`
	IsOwner bool        `json:"isOwner"`
	Conf    string      `json:"conf"`
}

func (createBank *CreateBankRequest) SetDefaults() {
	if createBank.State == state.Undefined {
		createBank.State = state.Available
	}
}

type GetBankRequest struct {
	MSPId   string `json:"mspId" validate:"required"`
	Address string `json:"address" validate:"required,validHex40"`
}

type SignDto struct {
	R string `json:"r" validate:"required,validHex64"`
	S string `json:"s" validate:"required,validHex64"`
	V int    `json:"v" validate:"required,gte=27,lte=28"`
}
