package requests

type TopupRequest struct {
	AddressTo     string `json:"addressTo" valid:"required~60302,validHex40~60301"`
	Amount        int64  `json:"amount" valid:"required~60313"`
	CurrencyCode  int    `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
	Payload       string `json:"payload"`
	TransactionId string `json:"transactionId" valid:"required~60331,uuidv4"`
}
