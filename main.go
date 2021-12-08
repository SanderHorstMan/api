package main

import (
	"fmt"

	"github.com/SanderIsdeManInhetZand/golangapi/database"
	"github.com/SanderIsdeManInhetZand/golangapi/magazijn"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloApi(c *fiber.Ctx) error {

	return c.SendString("Hello, MagaZijnAPI!")

}

func setupRoutes(app *fiber.App) {

	app.Get("/api/v1/product", magazijn.GetProducts)
	app.Get("/api/v1/product/:id", magazijn.GetProduct)
	app.Post("/api/v1/product", magazijn.NewProduct)
	app.Delete("/api/v1/product/:id", magazijn.DelteProduct)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "products.db")
	if err != nil {
		panic("niet geconnect met de db")

	}
	fmt.Println("succesvol geconnect met de database")
	database.DBConn.AutoMigrate(&magazijn.Product{})
	fmt.Println("Database gemigreerd")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	app.Get("/", helloApi)

	setupRoutes(app)

	app.Listen(":3000")
}
