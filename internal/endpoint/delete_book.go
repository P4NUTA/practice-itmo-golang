package endpoint

import (
	"github.com/gofiber/fiber/v2"
)

// DeleteBook godoc
//
//		@Summary		Delete book
//		@Description	Delete book
//		@Tags			books
//		@Accept			json
//		@Produce		json
//	 	@Param			book body BookRequest false "book info"
//		@Success		200	{object}	bool
//		@Security		ApiKeyAuth
//		@Router			/delete-book [post]
func (c *Controller) DeleteBook(ctx *fiber.Ctx) error {
	req := new(BookRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	id, err := c.Book.DeleteBook(ctx, req.ID)
	if err != nil {
		return err
	}

	return ctx.JSON(id)
}
