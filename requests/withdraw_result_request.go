package requests

type WithdrawResultRequest struct {
	AddressFrom string `json:"addressFrom" validate:"required,validHex40"`
	TxId        string `json:"txId" validate:"required"`
}
