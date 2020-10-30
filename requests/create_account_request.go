package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/account-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/identity-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
)

type CreateAccountRequest struct {
	Address       string                            `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	State         state_enum.State                  `json:"state"`
	CurrencyCode  int                               `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)"`
	JuridicalType juridical_type_enum.JuridicalType `json:"juridicalType"`
	IdentityType  identity_type_enum.IdentityType   `json:"identityType"`
	Type          account_type_enum.AccountType     `json:"type"`
	Identifiers   []string                          `json:"identifiers" valid:"validHex64~ErrorIdentifierNotFolowingRegex"`
	PublicKey     string                            `json:"publicKey"`
	Owner         string                            `json:"owner"`
	MsgHash       string                            `json:"msgHash"`
	Sig           SignDto                           `json:"sig"`
}

func (createAccount *CreateAccountRequest) SetDefaults() {
	if createAccount.IdentityType == identity_type_enum.Undefined {
		createAccount.IdentityType = identity_type_enum.None
	}
	if createAccount.State == state_enum.Undefined {
		createAccount.State = state_enum.Available
	}
	if createAccount.Type == account_type_enum.Undefined {
		createAccount.Type = account_type_enum.Client
	}
	if createAccount.Identifiers == nil {
		createAccount.Identifiers = []string{}
	}
}
