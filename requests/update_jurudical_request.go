package requests

import juridical_type_enum "github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type-enum"

type UpdateJuridicalRequest struct {
	Address       string                            `json:"address" valid:"required~60302,validHex40~60301"`
	JuridicalType juridical_type_enum.JuridicalType `json:"juridicalType" valid:"required~60345"`
}
