package invoice_state_enum

type InvoiceState int

const (
	Undefined InvoiceState = iota
	Created
	Billed
	Failed
	Expired
	Paid
	ClarificationRequested
)

func (invoiceState InvoiceState) Str() string {
	return [...]string{"Undefined", "Created", "Billed", "Failed", "Expired", "Paid", "ClarificationRequested"}[invoiceState]
}
