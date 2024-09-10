package loan

import (
	"amartha-test/internals/models"
	"context"

	"github.com/google/uuid"
)

type LoanAccountRepository interface {
	Insert(ctx context.Context, lam models.LoanAccount) error
	FindAll(ctx context.Context) ([]models.LoanAccount, error)
	Update(ctx context.Context, lam models.LoanAccount) (models.LoanAccount, error)
	GetLoanAccountById(ctx context.Context, id uuid.UUID) (models.LoanAccount, error)
	UpdateProposedLoan(ctx context.Context, lam models.LoanAccount) (models.LoanAccount, error)
	UpdateLoanDoc(ctx context.Context, lam models.LoanAccount) (models.LoanAccount, error)
	UpdateLoanInvested(ctx context.Context, lam models.LoanAccount) (models.LoanAccount, error)
}

type LoanAccountService interface {
	Create(ctx context.Context, lam models.LoanAccount) (string, error)
	FindAll(ctx context.Context) ([]models.LoanAccount, error)
	Update(ctx context.Context, cm models.LoanAccount) (models.LoanAccount, error)
	GetLoanAccountById(ctx context.Context, id uuid.UUID) (models.LoanAccount, error)
	UpdateProposedLoan(ctx context.Context, lam models.LoanAccount) (models.LoanAccount, error)
	UploadDocument(ctx context.Context, doc models.Document, id uuid.UUID) (models.LoanAccount, error)
	LoanInvest(ctx context.Context, li models.LoanInvested, id uuid.UUID) (models.LoanAccount, error)
}
