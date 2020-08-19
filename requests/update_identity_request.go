package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/identity-type-enum"

type UpdateIdentityRequest struct {
	Address      string                          `json:"address" valid:"required,validHex40"`
	IdentityType identity_type_enum.IdentityType `json:"identityType" valid:"required"`
}
