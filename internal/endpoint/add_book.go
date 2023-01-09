package endpoint

import (
	"github.com/gofiber/fiber/v2"
)

// AddBook godoc
//
//		@Summary		Add book
//		@Description	Add new book
//		@Tags			books
//		@Accept			json
//		@Produce		json
//	 	@Param			book body AddBookRequest false "book info"
//		@Success		200	{object}	int
//		@Security		ApiKeyAuth
//		@Router			/add-book [post]
func (c *Controller) AddBook(ctx *fiber.Ctx) error {
	req := new(AddBookRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	id, err := c.Book.AddBook(ctx, req.Name, req.Author, req.Description)
	if err != nil {
		return err
	}

	return ctx.JSON(id)
}
