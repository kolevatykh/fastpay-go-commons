package requests

type TransferRequest struct {
	AddressFrom   string  `json:"addressFrom" valid:"required~60302,validHex40~60301"`
	To            string  `json:"to" valid:"required~60302,validHex40or64~60367"`
	CurrencyCode  int     `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
	Amount        int64   `json:"amount" valid:"required~60313"`
	Payload       string  `json:"payload"`
	MsgHash       string  `json:"msgHash"`
	Sig           SignDto `json:"sig"`
	InvoiceNumber string  `json:"invoiceNumber" valid:"maxstringlength(255)~60353"`
	Exp           int64   `json:"exp~60332"`
	TransactionId string  `json:"transactionId" valid:"required~60331,uuidv4"`
}
