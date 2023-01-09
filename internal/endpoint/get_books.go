package endpoint

import "github.com/gofiber/fiber/v2"

// GetBooks godoc
//
//	@Summary		Get book list
//	@Description	get book list
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	[]BookResponse
//	@Security		ApiKeyAuth
//	@Router			/get-books [post]
func (c *Controller) GetBooks(ctx *fiber.Ctx) error {
	books, err := c.Book.GetBooks(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(books)
}
