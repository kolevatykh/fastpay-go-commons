package requests

type TransfersFromRequest struct {
	AddressFrom   string            `json:"addressFrom" valid:"required,validHex40"`
	To            []TransfersFromTo `json:"to" valid:"required"`
	CurrencyCode  int               `json:"currencyCode" valid:"required,min(0)"`
	TransactionId string            `json:"transactionId" valid:"required,uuidv4"`
}

type TransfersFromTo struct {
	To      string `json:"to" valid:"required,validHex40or64"`
	Amount  int64  `json:"amount" valid:"required"`
	Payload string `json:"payload"`
}
