package cservice

import (
	"amartha-test/internals/components"
	"amartha-test/internals/components/customer"
)

type customerService struct {
	repo components.Repository
}

var _ customer.CustomerService = (*customerService)(nil)

func NewCustomerService(
	repo components.Repository,
) *customerService {
	return &customerService{
		repo: repo,
	}
}
