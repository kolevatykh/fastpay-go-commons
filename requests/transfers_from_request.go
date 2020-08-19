package requests

type TransfersFromRequest struct {
	AddressFrom   string            `json:"addressFrom" valid:"required~60302,validHex40~60301"`
	To            []TransfersFromTo `json:"to" valid:"required"`
	CurrencyCode  int               `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
	TransactionId string            `json:"transactionId" valid:"required~60331,uuidv4"`
}

type TransfersFromTo struct {
	To      string `json:"to" valid:"required~60302,validHex40or64~60367"`
	Amount  int64  `json:"amount" valid:"required~60313"`
	Payload string `json:"payload"`
}
