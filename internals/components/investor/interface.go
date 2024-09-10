package investor

import (
	"amartha-test/internals/models"
	"context"

	"github.com/google/uuid"
)

type InvestorRepository interface {
	GetInvestorById(ctx context.Context, id uuid.UUID) (models.Investor, error)
}
