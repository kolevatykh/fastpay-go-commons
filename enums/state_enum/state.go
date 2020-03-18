package state_enum

type State int

const (
	Undefined State = iota
	Available
	Blocked
	TempBlocked
)

func (accountState State) String() string {
	return [...]string{"Undefined", "Available", "Blocked", "TempBlocked"}[accountState]
}
