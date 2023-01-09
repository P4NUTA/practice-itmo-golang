package service

import (
	"fmt"
	"practice/internal/models"
	"practice/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type service struct {
	repository repository.Repository
}

type Service interface {
	GetBook(ctx *fiber.Ctx, id int64) (*models.Book, error)
	GetBooks(ctx *fiber.Ctx) ([]models.Book, error)
	AddBook(ctx *fiber.Ctx, name, author, description string) (int64, error)
	DeleteBook(ctx *fiber.Ctx, id int64) (bool, error)
}

func InitService() Service {
	return &service{
		repository: repository.InitRepository(),
	}
}

func (s *service) GetBook(ctx *fiber.Ctx, id int64) (*models.Book, error) {
	return s.repository.GetBook(ctx, id)
}

func (s *service) GetBooks(ctx *fiber.Ctx) ([]models.Book, error) {
	return s.repository.GetBooks(ctx)
}

func (s *service) AddBook(ctx *fiber.Ctx, name, author, description string) (int64, error) {
	if name == "" {
		return 0, fmt.Errorf("name is empty")
	}

	return s.repository.AddBook(ctx, name, author, description)
}

func (s *service) DeleteBook(ctx *fiber.Ctx, id int64) (bool, error) {
	return s.repository.DeleteBook(ctx, id)
}
