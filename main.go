package main

import (
	"github.com/gofiber/fiber"
	"github.com/mazhaboy/test/tree/master/organization"
)

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}
func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/organizations", organization.GetOrganizations)
	app.Get("/organizations/:id", organization.GetOrganizationByID)
	app.Post("/organizations", organization.NewOrganization)
	app.Delete("/organizations/:id", organization.DeleteOrganization)
	app.Put("/organizations/:id", organization.UpdateOrganization)
}
func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}
