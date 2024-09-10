package laservice

import (
	"amartha-test/internals/models"
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

func (las *loanAccountService) Create(ctx context.Context, lam models.LoanAccount) (loanIDNo string, err error) {
	cr := las.repo.GetLoanAccountRepository()

	loanId := uuid.New()
	lam.ID = loanId
	loanIDNo = loanId.String()
	lam.Status = "pending"
	lam.LoanState = "proposed"
	lam.CreateBy = "system"
	lam.CreateDate = time.Now()

	err = cr.Insert(ctx, lam)
	if err != nil {
		slog.Error("Error create loan account:", err)
		return loanIDNo, err
	}

	return loanIDNo, err
}
