package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
	"practice/internal/models"
)

type Repository interface {
	GetBook(ctx *fiber.Ctx, id int64) (*models.Book, error)
	GetBooks(ctx *fiber.Ctx) ([]models.Book, error)
	AddBook(ctx *fiber.Ctx, name, author, description string) (int64, error)
	DeleteBook(ctx *fiber.Ctx, id int64) (bool, error)
}

type Client struct {
	db *pg.DB
}

func InitRepository() *Client {
	return &Client{
		db: pg.Connect(&pg.Options{
			User:     "postgres",
			Password: "postgres",
			Database: "postgres",
			Addr:     "localhost:5432",
		}),
	}
}

func (c Client) GetBook(ctx *fiber.Ctx, id int64) (*models.Book, error) {
	book := &models.Book{}
	if err := c.db.Model(book).Where("id = ?", id).Select(); err != nil {
		return nil, err
	}

	return book, nil
}

func (c Client) GetBooks(ctx *fiber.Ctx) ([]models.Book, error) {
	books := make([]models.Book, 0)
	if err := c.db.Model(&books).Select(); err != nil {
		return nil, err
	}

	return books, nil
}

func (c Client) AddBook(ctx *fiber.Ctx, name, author, description string) (int64, error) {
	book := &models.Book{
		Name:        name,
		Author:      author,
		Description: description,
	}

	if _, err := c.db.Model(book).Insert(); err != nil {
		return 0, err
	}
	return book.ID, nil
}

func (c Client) DeleteBook(ctx *fiber.Ctx, id int64) (bool, error) {
	if _, err := c.db.Model(&models.Book{}).Where("id = ?", id).Delete(); err != nil {
		return false, err
	}
	return true, nil
}
