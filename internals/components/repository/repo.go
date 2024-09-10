package repository

import (
	"gorm.io/gorm"

	"amartha-test/internals/components"
	"amartha-test/internals/components/customer"
	"amartha-test/internals/components/investor"
	"amartha-test/internals/components/loan"
)

var _ components.Repository = (*repository)(nil)

type dbrepo struct {
	r *repository
}

type repository struct {
	dbs       *gorm.DB
	customerR *customerRepo
	loanAr    *loanAccountRepo
	investorR *investorRepo
	//lir    *loanInvestorRepo
	common dbrepo
}

func NewRepository(db *gorm.DB) *repository {
	rtx := &repository{
		dbs: db,
	}
	rtx.common.r = rtx
	rtx.customerR = (*customerRepo)(&rtx.common)
	rtx.loanAr = (*loanAccountRepo)(&rtx.common)
	rtx.investorR = (*investorRepo)(&rtx.common)
	return rtx
}

func (r *repository) GetCustomerRepository() customer.CustomerRepository {
	return r.customerR
}

func (r *repository) GetLoanAccountRepository() loan.LoanAccountRepository {
	return r.loanAr
}

func (r *repository) GetInvestorRepository() investor.InvestorRepository {
	return r.investorR
}
