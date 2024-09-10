package laservice

import (
	"amartha-test/internals/models"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"golang.org/x/exp/slog"
)

func (las *loanAccountService) LoanInvest(ctx context.Context, li models.LoanInvested, id uuid.UUID) (lam models.LoanAccount, err error) {
	lar := las.repo.GetLoanAccountRepository()

	var loanAccountDb models.LoanAccount
	//var loanInvestedDb models.LoanInvested
	var loanInvested []models.LoanInvested
	loanInvestId := uuid.New()
	li.ID = loanInvestId
	li.LoanID = id
	li.CreateAt = time.Now()
	li.CreateBy = "system"

	loanAccountDb, err = lar.GetLoanAccountById(ctx, id)
	if err != nil {
		slog.Error("error update loan invest: ", err)
		return lam, err
	}

	loanInvested = loanAccountDb.Investors
	totalLoanInvested, err := totalInvested(loanInvested)
	if err != nil {
		slog.Debug("failed to total invest amount", err)
		err = errors.New("failed to total invest amount")
		return lam, err
	}

	if totalLoanInvested.GreaterThan(loanAccountDb.Principal) {
		err = errors.New("total Invest more than principal")
		return loanAccountDb, err
	} else if totalLoanInvested.Equal(loanAccountDb.Principal) {
		err = errors.New("loan is fully invested")
		return loanAccountDb, err
	}

	loanInvested = append(loanInvested, li)
	loanAccountDb.Investors = loanInvested

	lam, err = lar.UpdateLoanInvested(ctx, loanAccountDb)
	if err != nil {
		slog.Error("Error update loan invest data: ", err)
		return lam, err
	}

	lam, err = lar.GetLoanAccountById(ctx, id)
	if err != nil {
		slog.Error("Error update loan account: ", err)
		return lam, err
	}

	return
}

func totalInvested(lia []models.LoanInvested) (total decimal.Decimal, err error) {
	if len(lia) == 0 {
		total = decimal.NewFromInt(0)
		return
	}

	var sum decimal.Decimal
	for _, loansInvest := range lia {
		sum = sum.Add(loansInvest.InvestAmount)
	}

	return sum, err
}
