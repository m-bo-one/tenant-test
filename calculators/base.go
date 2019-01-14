package calculators

type Bonus interface {
	GetValue() float64
	SetValue(value float64)
	CalculateTotal() float64
}
