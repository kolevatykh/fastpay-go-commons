package filter_contract_state_enum

type FilterContractState int

const (
	All FilterContractState = iota
	Active
	Terminated
)

func (filterContractState FilterContractState) String() string {
	return [...]string{"All", "Active", "Terminated"}[filterContractState]
}
