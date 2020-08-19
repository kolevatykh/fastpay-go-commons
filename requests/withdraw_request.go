package requests

type WithdrawRequest struct {
	AddressFrom   string  `json:"addressFrom" valid:"required,validHex40"`
	Amount        int64   `json:"amount" valid:"required"`
	CurrencyCode  int     `json:"currencyCode" valid:"required,min(0)"`
	Payload       string  `json:"payload"`
	TransactionId string  `json:"transactionId" valid:"required,uuidv4"`
	MsgHash       string  `json:"msgHash" valid:"required"`
	Exp           int64   `json:"exp" valid:"required"`
	Sig           SignDto `json:"sig" valid:"required"`
}
