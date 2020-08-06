package deal_state_enum

type DealState int

const (
	Undefined            DealState = 0
	ApplicationCreated   DealState = 1
	ApplicationPublished DealState = 2
	ContractAwarded      DealState = 3
	ContractTerminated   DealState = 4
	ContractExecuted     DealState = 5
)

func (state DealState) String() string {
	return [...]string{"Undefined", "ApplicationCreated", "ApplicationPublished", "ContractAwarded", "ContractTerminated", "ContractExecuted"}[state]
}
