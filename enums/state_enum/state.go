package state_enum

type State int

const (
	Undefined State = iota
	Available
	Blocked
	TempBlocked
	Deleted
)

func (accountState State) Str() string {
	return [...]string{"Undefined", "Available", "Blocked", "TempBlocked", "Deleted"}[accountState]
}
