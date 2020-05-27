package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/identity-type-enum"

type GetAccountLimitsRequest struct {
	IdentityType   identity_type_enum.IdentityType `json:"identityType" validate:"required,gte=0"`
	CurrencyCode   int                             `json:"currencyCode" validate:"required,gte=0,lte=999"`
	AccountAddress string                          `json:"accountAddress" validate:"required,validHex40"`
}
