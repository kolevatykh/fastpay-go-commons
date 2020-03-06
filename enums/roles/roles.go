package roles

type Roles int

const (
	Undefined Roles = iota
	Owner
	Bank
	Regulator
	Any
)

func (role Roles) String() string {
	return [...]string{"Undefined", "Owner", "Bank", "Regulator", "Any"}[role]
}
