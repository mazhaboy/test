package organization

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/mazhaboy/test/tree/master/database"
)

type Organization struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func GetOrganizations(c *fiber.Ctx) {
	db := database.DBConn
	var organizations []Organization
	db.Find(&organizations)
	c.JSON(organizations)
}
func GetOrganizationByID(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var organization Organization
	db.Find(&organization, id)
	c.JSON(organization)
}
func NewOrganization(c *fiber.Ctx) {
	db := database.DBConn
	var organization Organization
	organization.Name = "Dar"
	organization.Address = "Abay 10"
	organization.Phone = "+7023640050"
	db.Create(&organization)
	c.JSON(organization)
}
func DeleteOrganization(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var organization Organization
	db.First(&organization, id)
	if organization.Name == "" {
		c.Status(500).Send("No organization found with ID")
		return
	}
	db.Delete(&organization)
	c.Send("Organization successfully deleted")
}
func UpdateOrganization(c *fiber.Ctx) {
	c.Send("Update Organization")
}
