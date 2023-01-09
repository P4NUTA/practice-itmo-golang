package endpoint

import "github.com/gofiber/fiber/v2"

// Main godoc
//
//	@Summary		test connection
//	@Description	test connection
//	@Tags			test
//	@Accept			json
//	@Produce		plain
//	@Success		200	{object}	string
//	@Security		ApiKeyAuth
//	@Router			/main [get]
func (c *Controller) Main(ctx *fiber.Ctx) error {
	return ctx.SendString("HELLO WORLD!")
}
