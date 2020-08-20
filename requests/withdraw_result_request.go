package requests

type WithdrawResultRequest struct {
	AddressFrom string `json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	TxId        string `json:"txId" valid:"required~ErrorTxIdSNotPassed"`
}
