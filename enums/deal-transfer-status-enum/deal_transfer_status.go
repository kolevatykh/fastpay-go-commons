package deal_transfer_status_enum

type DealTransferStatus int

const (
	Undefined DealTransferStatus = iota
	Awaiting
	Canceled
	Performed
)

func (dealTransferStatus DealTransferStatus) Str() string {
	return [...]string{"Undefined", "Awaiting", "Canceled", "Performed"}[dealTransferStatus]
}

func Parse(dealTransferStatus string) DealTransferStatus {
	switch dealTransferStatus {
	case "Awaiting":
		return Awaiting
	case "Canceled":
		return Canceled
	case "Performed":
		return Performed
	default:
		return Undefined
	}
}

func GetAllDealTransferStatuses() []DealTransferStatus {
	return []DealTransferStatus{Undefined, Awaiting, Canceled, Performed}
}
