package calculators

type DefaultBonus struct {
	value float64
}

type SuperBonus struct {
	value float64
}

func (b *DefaultBonus) GetValue() float64 {
	return b.value
}

func (b *DefaultBonus) SetValue(value float64) {
	b.value = value
}

func (b *DefaultBonus) CalculateTotal() float64 {
	return b.GetValue() * 0.45
}

func (b *SuperBonus) GetValue() float64 {
	return b.value
}

func (b *SuperBonus) SetValue(value float64) {
	b.value = value
}

func (b *SuperBonus) CalculateTotal() float64 {
	return b.GetValue() * 0.85
}
