package requests

type WithdrawResultRequest struct {
	AddressFrom string `json:"addressFrom" valid:"required~60302,validHex40~60301"`
	TxId        string `json:"txId" valid:"required~60331"`
}
