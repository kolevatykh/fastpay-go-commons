package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/identity-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/limit-type-enum"
)

type GetLimitRequest struct {
	LimitType     limit_type_enum.LimitType         `json:"limitType" valid:"required,gte(0)"`
	IdentityType  identity_type_enum.IdentityType   `json:"identityType" valid:"required,gte(0)"`
	JuridicalType juridical_type_enum.JuridicalType `json:"juridicalType" valid:"gte(0)"`
	CurrencyCode  int                               `json:"currencyCode" valid:"required,gte(0),lte(999)"`
}
