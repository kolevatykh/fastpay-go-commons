package requests

import identity_type_enum "github.com/SolarLabRU/fastpay-go-commons/enums/identity-type-enum"

type UpdateIdentityRequest struct {
	Address      string                          `json:"address" validate:"required,validHex40"`
	IdentityType identity_type_enum.IdentityType `json:"identityType" validate:"required"`
}
