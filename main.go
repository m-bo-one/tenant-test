package main

import (
	"github.com/tenant-test/calculators"
	"github.com/tenant-test/controllers"
	"github.com/tenant-test/models"

	"github.com/kataras/iris"
)

func main() {
	defaultSchema := &models.Schema{
		Bonus: calculators.Default,
	}
	customSchema := &models.Schema{
		Bonus: calculators.Super,
	}
	tenants := []models.Tenant{
		models.Tenant{Host: "hello.local", Schema: defaultSchema},
		models.Tenant{Host: "hello2.local", Schema: customSchema},
	}
	models.InitTenants(models.TenantStore, tenants)

	app := iris.Default()
	app.Get("/bonuses", controllers.GetBonus)
	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":80"))
}
