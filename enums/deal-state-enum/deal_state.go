package deal_state_enum

type DealState int

const (
	Undefined                       DealState = 0
	ApplicationCreated              DealState = 1
	ApplicationDismissedByInitiator DealState = 2
	ApplicationRejected             DealState = 3
	ApplicationPublished            DealState = 4
	ContractAwarded                 DealState = 5
	ContractTerminated              DealState = 6
	ContractExecuted                DealState = 7
)

func (state DealState) String() string {
	return [...]string{"Undefined", "ApplicationCreated", "ApplicationDismissedByInitiator", "ApplicationRejected", "ApplicationPublished", "ContractAwarded", "ContractTerminated", "ContractExecuted"}[state]
}
