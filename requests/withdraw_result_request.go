package requests

type WithdrawResultRequest struct {
	AddressFrom string `json:"addressFrom" valid:"required,validHex40"`
	TxId        string `json:"txId" valid:"required"`
}
