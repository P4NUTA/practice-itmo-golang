package main

import (
	"practice/internal/endpoint"

	"github.com/gofiber/fiber/v2"
)

// @title			Fiber Example API
// @version		1.0
// @description	This is a sample swagger for Fiber
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	fiber@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/
func main() {
	app := fiber.New()

	endpoint.SetupRoutes(app)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
