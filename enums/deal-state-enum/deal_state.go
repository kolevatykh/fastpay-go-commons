package deal_state_enum

type DealState int

const (
	Undefined            DealState = 0
	ApplicationCreated   DealState = 1
	ApplicationDismissed DealState = 2
	ApplicationRejected  DealState = 3
	ApplicationPublished DealState = 4
	ContractAwarded      DealState = 5
	ContractTerminated   DealState = 6
	ContractExecuted     DealState = 7
	ContractDisputeOpen  DealState = 8
)

func (state DealState) Str() string {
	return [...]string{"Undefined", "ApplicationCreated", "ApplicationDismissed", "ApplicationRejected", "ApplicationPublished", "ContractAwarded", "ContractTerminated", "ContractExecuted", "ContractDisputeOpen"}[state]
}
