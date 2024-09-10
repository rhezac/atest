package cservice

import (
	"amartha-test/internals/models"
	"context"
	"time"

	"golang.org/x/exp/slog"
)

func (cs *customerService) Update(ctx context.Context, cm models.Customer) (cust models.Customer, err error) {
	cr := cs.repo.GetCustomerRepository()

	cm.UpdateBy = "system"
	cm.UpdateDate = time.Now()

	cust, err = cr.Update(ctx, cm)
	if err != nil {
		slog.Error("Error update customer: ", err)
		return cust, err
	}
	return cust, err
}
