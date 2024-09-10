package cservice

import (
	"amartha-test/internals/models"
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

func (cs *customerService) Create(ctx context.Context, cm models.Customer) (custIDNo string, err error) {
	cr := cs.repo.GetCustomerRepository()

	custId := uuid.New()
	cm.ID = custId
	custIDNo = custId.String()

	cm.CreateBy = "system"
	cm.CreateDate = time.Now()

	err = cr.Insert(ctx, cm)
	if err != nil {
		slog.Error("Error create customer:", err)
		return custIDNo, err
	}

	return custIDNo, err
}
