package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/identity-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/limit-type-enum"
)

type GetLimitRequest struct {
	LimitType     limit_type_enum.LimitType         `json:"limitType" valid:"required,range(0|9223372036854775807)"`
	IdentityType  identity_type_enum.IdentityType   `json:"identityType" valid:"required,range(0|9223372036854775807)~60305"`
	JuridicalType juridical_type_enum.JuridicalType `json:"juridicalType" valid:"range(0|9223372036854775807)~60307"`
	CurrencyCode  int                               `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
}
