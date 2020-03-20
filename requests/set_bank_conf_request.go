package requests

type SetBankConfRequest struct {
	Conf string `json:"conf" validate:"required"`
}
