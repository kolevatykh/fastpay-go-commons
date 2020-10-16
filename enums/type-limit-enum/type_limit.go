package type_limit_enum

type TypeLimit uint

const (
	Daily TypeLimit = iota
	Monthly
)

func (typeLimit TypeLimit) Str() string {
	return [...]string{"Daily", "Monthly"}[typeLimit]
}
