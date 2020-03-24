package requests

type GetAccountBalanceRequest struct {
	Address string `json:"address" validate:"required,validHex40"`
}
