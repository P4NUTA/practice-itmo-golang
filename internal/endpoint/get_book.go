package endpoint

import (
	"github.com/gofiber/fiber/v2"
)

// GetBook godoc
//
//		@Summary		Get book
//		@Description	get book info
//		@Tags			books
//		@Accept			json
//		@Produce		json
//	 	@Param			book body BookRequest true "book id"
//		@Success		200	{object}	BookResponse
//		@Security		ApiKeyAuth
//		@Router			/get-book [post]
func (c *Controller) GetBook(ctx *fiber.Ctx) error {
	req := new(BookRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	book, err := c.Book.GetBook(ctx, req.ID)
	if err != nil {
		return err
	}

	return ctx.JSON(book)
}
