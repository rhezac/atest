package repository

import (
	"amartha-test/internals/components/loan"
	"amartha-test/internals/models"
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type loanAccountRepo dbrepo

var _ loan.LoanAccountRepository = (*loanAccountRepo)(nil)

func (lar *loanAccountRepo) Insert(ctx context.Context, lam models.LoanAccount) (err error) {
	err = lar.r.dbs.Create(&lam).Error
	if err != nil {
		slog.Error("Failed to insert loan account data to db")
		return err
	}

	return
}

func (lar *loanAccountRepo) FindAll(ctx context.Context) (loans []models.LoanAccount, err error) {
	err = lar.r.dbs.Find(&loans).Error
	if err != nil {
		slog.Error("Failed to find loan account data from db")
		return loans, err
	}

	return loans, err
}

func (lar *loanAccountRepo) Update(ctx context.Context, lam models.LoanAccount) (ln models.LoanAccount, err error) {
	err = lar.r.dbs.Where("id = ?", lam.ID).Save(lam).Error
	if err != nil {
		slog.Error("Failed to update loan account data to db")
		return lam, err
	}

	return lam, err
}

func (lar *loanAccountRepo) UpdateProposedLoan(ctx context.Context, lam models.LoanAccount) (ln models.LoanAccount, err error) {
	err = lar.r.dbs.Model(&lam).Updates(models.LoanAccount{Rate: lam.Rate, Principal: lam.Principal, ROI: lam.ROI, TenorValue: lam.TenorValue}).Error
	if err != nil {
		slog.Error("Failed to update loan account data to db")
		return lam, err
	}
	return lam, err
}

func (lar *loanAccountRepo) UpdateLoanDoc(ctx context.Context, lam models.LoanAccount) (ln models.LoanAccount, err error) {
	loanDocs, err := json.Marshal(lam.Documents)
	if err != nil {
		slog.Error("Failed to convert documents to string")
	}

	err = lar.r.dbs.Model(&lam).Update("documents", string(loanDocs)).Error
	if err != nil {
		slog.Error("Failed to update loan account document data to db")
		return ln, err
	}
	return lam, err
}

func (lar *loanAccountRepo) UpdateLoanInvested(ctx context.Context, lam models.LoanAccount) (ln models.LoanAccount, err error) {
	loanInvest, err := json.Marshal(lam.Investors)
	if err != nil {
		slog.Error("Failed to convert loan invest to string")
	}

	err = lar.r.dbs.Model(&lam).Update("investors", string(loanInvest)).Error
	if err != nil {
		slog.Error("Failed to update loan account investors data to db")
		return ln, err
	}

	return lam, err
}

func (lar *loanAccountRepo) GetLoanAccountById(ctx context.Context, id uuid.UUID) (ln models.LoanAccount, err error) {
	err = lar.r.dbs.Where("id = ?", id).First(&ln).Error
	if err != nil {
		slog.Error("Failed to find loan account by id from db")
		return ln, err
	}

	return ln, err
}
