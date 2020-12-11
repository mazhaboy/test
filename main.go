package main

import (
	"github.com/gofiber/fiber"
	"github.com/mazhaboy/test/organization"
)

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}
func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/organizations", organization.GetOrganizations)
	app.Get("/organizations/:id", helloWorld)
	app.Post("/organizations", helloWorld)
	app.Delete("/organizations/:id", helloWorld)
	app.Put("/organizations/:id", helloWorld)
}
func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}
