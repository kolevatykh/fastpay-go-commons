package requests

type SetBankConfRequest struct {
	TechnicalSignRequest
	Conf string `json:"conf" valid:"required"`
}
