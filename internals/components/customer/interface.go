package customer

import (
	"amartha-test/internals/models"
	"context"

	"github.com/google/uuid"
)

type CustomerRepository interface {
	Insert(ctx context.Context, cm models.Customer) error
	FindAll(ctx context.Context) ([]models.Customer, error)
	Update(ctx context.Context, cm models.Customer) (models.Customer, error)
	GetCustomerById(ctx context.Context, id uuid.UUID) (models.Customer, error)
}

type CustomerService interface {
	Create(ctx context.Context, cm models.Customer) (string, error)
	FindAll(ctx context.Context) ([]models.Customer, error)
	Update(ctx context.Context, cm models.Customer) (models.Customer, error)
	GetCustomerById(ctx context.Context, id uuid.UUID) (models.Customer, error)
}
