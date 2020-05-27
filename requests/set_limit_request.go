package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/identity-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/limit-type-enum"
)

type SetLimitRequest struct {
	LimitType    limit_type_enum.LimitType       `json:"limitType" validate:"required,gte=0"`
	IdentityType identity_type_enum.IdentityType `json:"identityType" validate:"required,gte=0"`
	Value        int64                           `json:"value" validate:"gte=0"`
	CurrencyCode int                             `json:"currencyCode" validate:"required,gte=0,lte=999"`
}
