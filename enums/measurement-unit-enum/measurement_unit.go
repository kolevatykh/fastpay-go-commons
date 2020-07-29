package measurement_unit_enum

type MeasurementUnit int

const (
	Undefined MeasurementUnit = 0
	// Литр
	Liter = 1
	// Килограмм
	Kilogram = 2
	// Фунт
	Pound = 3
	// Баррель
	Barrel = 4
)

func (measurementUnit MeasurementUnit) String() string {
	return [...]string{"Undefined", "Liter", "Kilogram", "Pound", "Barrel"}[measurementUnit]
}

func Parse(stringJuridicalType string) MeasurementUnit {
	switch stringJuridicalType {
	case "Liter":
		return Liter
	case "Kilogram":
		return Kilogram
	case "Pound":
		return Pound
	case "Barrel":
		return Barrel
	default:
		return Undefined
	}
}

func GetMeasurementUnits() []MeasurementUnit {
	return []MeasurementUnit{Undefined, Liter, Kilogram, Pound, Barrel}
}
