package magazijn

import (
	"github.com/SanderIsdeManInhetZand/golangapi/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Title       string `json:"title"`
	Leverancier string `json:"leverancier"`
	Nummer      int    `json:"nummer"`
}

func GetProducts(c *fiber.Ctx) error {
	db := database.DBConn
	var products []Product
	db.Find(&products)
	return c.JSON(products)

}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var product Product
	db.Find(&product, id)
	return c.JSON(product)

}
func NewProduct(c *fiber.Ctx) error {
	db := database.DBConn
	var product Product
	product.Title = "XMX100"
	product.Leverancier = "DHL"
	product.Nummer = 4321

	db.Create(&product)
	return c.JSON(product)

}
func DelteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var product Product
	db.First(&product, id)

	if product.Title == "" {

		return c.Status(500).SendString("Geen product gevonden met juiste ID")

	}
	db.Delete(&product)
	return c.SendString("Product verwijderd")

}
