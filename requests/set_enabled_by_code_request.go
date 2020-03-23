package requests

type SetEnabledByCodeRequest struct {
	Code    int  `json:"code" validate:"required,min=0"`
	Enabled bool `json:"enabled"`
}
