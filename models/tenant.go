package models

import (
	"github.com/tenant-test/calculators"
)

const DefaultBonus = 10

type Tenant struct {
	Host   string
	Schema *Schema
}

func (t *Tenant) GetBonusCalculator() (calc calculators.Bonus, err error) {
	calc = t.Schema.Bonus.GetSchema()
	calc.SetValue(DefaultBonus)
	return
}

func InitTenants(store map[string]*Tenant, tenants []Tenant) {
	for _, tenant := range tenants {
		func(t Tenant) {
			store[tenant.Host] = &t
		}(tenant)
	}
}
