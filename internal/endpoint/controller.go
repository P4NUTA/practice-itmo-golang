package endpoint

import (
	_ "practice/docs"
	"practice/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

const (
	get  = "GET"
	post = "POST"
)

type Controller struct {
	Book service.Service
}

func initController(book service.Service) Controller {
	return Controller{
		Book: book,
	}
}

type Route struct {
	Method  string
	Path    string
	Handler fiber.Handler
}

func (c *Controller) GetRoutes() []Route {
	return []Route{
		{
			Method:  get,
			Path:    "/main",
			Handler: c.Main,
		},
		{
			Method:  post,
			Path:    "/get-book",
			Handler: c.GetBook,
		},
		{
			Method:  post,
			Path:    "/add-book",
			Handler: c.AddBook,
		},
		{
			Method:  post,
			Path:    "/get-books",
			Handler: c.GetBooks,
		},
		{
			Method:  post,
			Path:    "/delete-book",
			Handler: c.DeleteBook,
		},
	}
}

func SetupRoutes(app *fiber.App) {

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		Next:             nil,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTION",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	c := initController(service.InitService())

	for _, route := range c.GetRoutes() {
		app.Add(route.Method, route.Path, route.Handler)
	}
}
