package repository

import (
	"amartha-test/internals/components/investor"
	"amartha-test/internals/models"
	"context"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type investorRepo dbrepo

var _ investor.InvestorRepository = (*investorRepo)(nil)

func (ir *investorRepo) GetInvestorById(ctx context.Context, id uuid.UUID) (investor models.Investor, err error) {
	err = ir.r.dbs.Where("id = ?", id).First(&investor).Error
	if err != nil {
		slog.Error("Failed to find customer by id from db")
		return investor, err
	}

	return investor, err
}
