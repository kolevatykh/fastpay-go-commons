package requests

type SetEnabledByCodeRequest struct {
	TechnicalSignRequest
	Code    int  `json:"code" valid:"required,range(0|9223372036854775807)"`
	Enabled bool `json:"enabled"`
}
