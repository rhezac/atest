package components

import (
	"amartha-test/internals/components/customer"
	"amartha-test/internals/components/investor"
	"amartha-test/internals/components/loan"
)

type Repository interface {
	GetCustomerRepository() customer.CustomerRepository
	GetLoanAccountRepository() loan.LoanAccountRepository
	GetInvestorRepository() investor.InvestorRepository
	//GetLoanInvestorRepository() LoanInvestorRepository
}
