package requests

import juridical_type_enum "github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type-enum"

type UpdateJuridicalRequest struct {
	Address       string                            `json:"address" validate:"required, validHex40"`
	JuridicalType juridical_type_enum.JuridicalType `json:"juridicalType" validate:"required"`
}
