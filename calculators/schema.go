package calculators

type BonusSchema int

func (schema BonusSchema) GetSchema() Bonus {
	cls := [...]Bonus{
		&DefaultBonus{},
		&SuperBonus{},
	}
	return cls[schema]
}

const (
	Default BonusSchema = 0
	Super   BonusSchema = 1
)
