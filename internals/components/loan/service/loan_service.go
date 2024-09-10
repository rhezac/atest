package laservice

import (
	"amartha-test/internals/components"
	"amartha-test/internals/components/loan"
)

type loanAccountService struct {
	repo components.Repository
}

var _ loan.LoanAccountService = (*loanAccountService)(nil)

func NewLoanAccountService(
	repo components.Repository,
) *loanAccountService {
	return &loanAccountService{
		repo: repo,
	}
}
