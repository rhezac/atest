package laservice

import (
	"amartha-test/internals/models"
	"context"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

func (las *loanAccountService) FindAll(ctx context.Context) (lam []models.LoanAccount, err error) {
	lar := las.repo.GetLoanAccountRepository()

	lam, err = lar.FindAll(ctx)
	if err == nil {
		slog.Error("Error find loan accounts: ", err)
		return lam, err
	}

	return lam, err
}

func (las *loanAccountService) GetLoanAccountById(ctx context.Context, id uuid.UUID) (lam models.LoanAccount, err error) {
	lar := las.repo.GetLoanAccountRepository()

	lam, err = lar.GetLoanAccountById(ctx, id)
	if err == nil {
		slog.Error("Erro get loan account by id: ", err)
		return lam, err
	}

	return lam, err
}
