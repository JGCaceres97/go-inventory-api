package service

import (
	"context"
	"errors"

	"github.com/jgcaceres97/go-inventory-api/src/internal/models"
)

var validRolesToAddProduct []int64 = []int64{1, 2}
var ErrInvalidPermission = errors.New("user does not have permission to add product")

func (s *serv) GetProducts(ctx context.Context) ([]models.Product, error) {
	productsEntity, err := s.repo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := []models.Product{}
	for _, product := range productsEntity {
		products = append(products, models.Product{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		})
	}

	return products, nil
}

func (s *serv) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	productEntity, err := s.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	product := &models.Product{
		ID:          productEntity.ID,
		Name:        productEntity.Name,
		Description: productEntity.Description,
		Price:       productEntity.Price,
	}

	return product, nil
}

func (s *serv) AddProduct(ctx context.Context, product models.Product, email string) error {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	roles, err := s.repo.GetUserRoles(ctx, user.ID)
	if err != nil {
		return err
	}

	userCanAdd := false

	for _, role := range roles {
		for _, validRole := range validRolesToAddProduct {
			if validRole == role.RoleID {
				userCanAdd = true
			}
		}
	}

	if !userCanAdd {
		return ErrInvalidPermission
	}

	return s.repo.SaveProduct(ctx, product.Name, product.Description, product.Price, user.ID)
}
