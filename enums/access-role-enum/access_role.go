package access_role_enum

type AccessRole int

const (
	Undefined AccessRole = 0
	Owner     AccessRole = 1 << 0
	Bank      AccessRole = 1 << 1
	Regulator AccessRole = 1 << 2
	OwnerShip AccessRole = 1 << 3
	Any                  = Owner | Bank | Regulator | OwnerShip
)

func (role AccessRole) Str() string {
	return [...]string{"Undefined", "Owner", "Bank", "Regulator", "OwnerShip", "Any"}[role]
}
