package requests

type SetLimitRequest struct {
	TechnicalSignRequest
	GetLimitRequest
	Value int64 `json:"value" valid:"range(0|9223372036854775807)"`
}
