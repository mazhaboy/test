package organization

import "github.com/gofiber/fiber"

func GetOrganizations(c *fiber.Ctx) {
	c.Send("All Organizations")
}
func GetOrganizationByID(c *fiber.Ctx) {
	c.Send("Single Organization")
}
func NewOrganization(c *fiber.Ctx) {
	c.Send("New Organization")
}
func DeleteOrganization(c *fiber.Ctx) {
	c.Send("Delete Organization")
}
func UpdateOrganization(c *fiber.Ctx) {
	c.Send("Update Organization")
}
