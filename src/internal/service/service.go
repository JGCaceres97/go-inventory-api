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
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{repo: repo}
}
