package controllers

import (
	"fmt"
	"github.com/tenant-test/models"

	"github.com/kataras/iris"
)

func GetBonus(ctx iris.Context) {
	request := ctx.Request()

	fmt.Printf("Resolve host: %v\n", request.Host)
	fmt.Printf("Current tenants: %+v\n", models.TenantStore)

	tenant, ok := models.TenantStore[request.Host]
	if !ok {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(ApiResource("ERROR", "Wrong host"))
		return
	}

	fmt.Printf("Found tenant: %v\n", tenant)

	calc, err := tenant.GetBonusCalculator()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(ApiResource("ERROR", err))
		return
	}

	fmt.Printf("Found calc: %+v\n", calc)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(ApiResource("OK", iris.Map{
		"total": calc.CalculateTotal(),
	}))
}
