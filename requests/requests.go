package requests

import (
	account_type "github.com/SolarLabRU/fastpay-go-commons/enums/account-type"
	identity_type "github.com/SolarLabRU/fastpay-go-commons/enums/identity-type"
	juridical_type "github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state"
)

type CreateAccount struct {
	Address       string                       `json:"address" validate:"required,validHex40"`
	State         state.State                  `json:"state" validate:"required"`
	CurrencyCode  int                          `json:"currencyCode" validate:"required,gte=0,lte=999"`
	JuridicalType juridical_type.JuridicalType `json:"juridicalType"`
	IdentityType  identity_type.IdentityType   `json:"identityType" validate:"required"`
	Type          account_type.AccountType     `json:"type" validate:"required"`
	Identifiers   []string                     `json:"identifiers" validate:"required,dive,validHex64"`
	Timestamp     int64                        `json:"timestamp" validate:"required"`
	PublicKey     string                       `json:"publicKey" validate:"required"`
	MsgHash       string                       `json:"msgHash" validate:"required"`
	Sig           SignDto                      `json:"sig" validate:"required,dive"`
}

type CreateBank struct {
	Address string      `json:"address"  validate:"required,validHex40"`
	State   state.State `json:"state" validate:"required"`
	Name    string      `json:"name"`
	BIK     string      `json:"bik"`
	MspId   string      `json:"mspId"`
	IsOwner bool        `json:"isOwner"`
	Conf    string      `json:"conf"`
}

type GetBank struct {
	MSPId   string `json:"mspId" validate:"required"`
	Address string `json:"address" validate:"required,validHex40"`
}

type SignDto struct {
	R string `json:"r" validate:"required,validHex64"`
	S string `json:"s" validate:"required,validHex64"`
	V int    `json:"v" validate:"required,gte=27,lte=28"`
}
