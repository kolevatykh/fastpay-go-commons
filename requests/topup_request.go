package requests

type TopupRequest struct {
	AddressTo     string `json:"addressTo" valid:"required,validHex40"`
	Amount        int64  `json:"amount" valid:"required"`
	CurrencyCode  int    `json:"currencyCode" valid:"required,min(0)"`
	Payload       string `json:"payload"`
	TransactionId string `json:"transactionId" valid:"required,uuidv4"`
}
