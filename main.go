package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/", helloWorld)
	app.Listen(3000)
}
func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}
