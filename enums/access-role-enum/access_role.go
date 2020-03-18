package access_role_enum

type AccessRole int

const (
	Undefined AccessRole = iota
	Owner
	Bank
	Regulator
	Any
)

func (role AccessRole) String() string {
	return [...]string{"Undefined", "Owner", "Bank", "Regulator", "Any"}[role]
}
