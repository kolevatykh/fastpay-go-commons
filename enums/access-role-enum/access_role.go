package access_role_enum

type AccessRole int

const (
	Undefined AccessRole = 0
	Owner     AccessRole = 1 << 0
	Bank      AccessRole = 1 << 1
	Regulator AccessRole = 1 << 2
	Any                  = Owner | Bank | Regulator
)

func (role AccessRole) Str() string {
	return [...]string{"Undefined", "Owner", "Bank", "Regulator", "Any"}[role]
}
