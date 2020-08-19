package requests

type SetEnabledByCodeRequest struct {
	Code    int  `json:"code" valid:"required,min(0)"`
	Enabled bool `json:"enabled"`
}
