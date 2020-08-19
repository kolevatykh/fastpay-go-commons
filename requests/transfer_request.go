package requests

type TransferRequest struct {
	AddressFrom   string  `json:"addressFrom" valid:"required,validHex40"`
	To            string  `json:"to" valid:"required,validHex40or64"`
	CurrencyCode  int     `json:"currencyCode" valid:"required,min(0)"`
	Amount        int64   `json:"amount" valid:"required"`
	Payload       string  `json:"payload"`
	MsgHash       string  `json:"msgHash"`
	Sig           SignDto `json:"sig" valid:""`
	InvoiceNumber string  `json:"invoiceNumber" valid:"max(255)"`
	Exp           int64   `json:"exp"`
	TransactionId string  `json:"transactionId" valid:"required,uuidv4"`
}
