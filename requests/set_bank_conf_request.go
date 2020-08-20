package requests

type SetBankConfRequest struct {
	Conf string `json:"conf" valid:"required"`
}
