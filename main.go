package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mazhaboy/test/tree/master/database"
	"github.com/mazhaboy/test/tree/master/organization"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "organizations.db")
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Successfully connected to the DB")
	database.DBConn.AutoMigrate(&organization.Organization{})
	fmt.Println("Database Migrated")
}

func main() {

	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()
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
