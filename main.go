package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mazhaboy/test/tree/master/database"
	"github.com/mazhaboy/test/tree/master/organization"
)

type Organization struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

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

	engine := html.NewFileSystem(http.Dir("./templates"), ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/organizations", GetOrganizations)
	app.Get("/organizations/:id", GetOrganizationByID)
	app.Post("/organizations", NewOrganization)
	app.Delete("/organizations/:id", DeleteOrganization)
	app.Put("/organizations/:id", UpdateOrganization)
	initDatabase()

	defer database.DBConn.Close()
	// setupRoutes(app)
	app.Listen(":3000")
}
func GetOrganizations(c *fiber.Ctx) error {
	db := database.DBConn
	var organizations []Organization
	db.Find(&organizations)
	return c.Render("article", fiber.Map{
		"Title": organizations,
	})

}
func GetOrganizationByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var organization Organization
	db.Find(&organization, id)
	return c.Render("index", fiber.Map{
		"Title": organization,
	})
}
func NewOrganization(c *fiber.Ctx) error {
	db := database.DBConn
	var organization Organization
	organization.Name = "Halyk"
	organization.Address = "Mustafina 4"
	organization.Phone = "+7029897653"
	db.Create(&organization)
	return c.Render("index", fiber.Map{
		"Title": organization,
	})
}
func DeleteOrganization(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var organization Organization
	db.First(&organization, id)
	if organization.Name == "" {
		return c.Render("article", fiber.Map{
			"Title": organization,
		})
	}
	db.Delete(&organization)
	return c.Render("article", fiber.Map{
		"Title": organization,
	})
}
func UpdateOrganization(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var organization Organization
	db.First(&organization, id)
	if organization.Name == "" {
		return c.Render("article", fiber.Map{
			"Title": organization,
		})
	}
	db.Delete(&organization)
	return c.Render("article", fiber.Map{
		"Title": organization,
	})
}

// func setupRoutes(app *fiber.App) {

// 	app.Get("/organizations", organization.GetOrganizations)
// 	app.Get("/organizations/:id", organization.GetOrganizationByID)
// 	app.Post("/organizations", organization.NewOrganization)
// 	app.Delete("/organizations/:id", organization.DeleteOrganization)
// 	app.Put("/organizations/:id", organization.UpdateOrganization)
// }
