package repository

import (
	"context"

	"github.com/jgcaceres97/go-inventory-api/src/internal/entity"
	"github.com/jmoiron/sqlx"
)

// Repository es la interfaz que engloba las operaciones CRUD.
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{db: db}
}
