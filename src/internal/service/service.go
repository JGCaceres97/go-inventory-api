package service

import (
	"context"

	"github.com/jgcaceres97/go-inventory-api/src/internal/models"
	"github.com/jgcaceres97/go-inventory-api/src/internal/repository"
)

// Service es la lógica de negocio de la aplicación.
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
	AddUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProduct(ctx context.Context, id int64) (*models.Product, error)
	AddProduct(ctx context.Context, product models.Product, email string) error
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{repo: repo}
}
