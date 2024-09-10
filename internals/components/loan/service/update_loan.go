package laservice

import (
	"amartha-test/internals/components/rules"
	"amartha-test/internals/models"
	"amartha-test/internals/utils"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

func (las *loanAccountService) Update(ctx context.Context, lam models.LoanAccount) (ln models.LoanAccount, err error) {
	var lamdb models.LoanAccount
	lar := las.repo.GetLoanAccountRepository()

	lamdb, err = lar.GetLoanAccountById(ctx, lam.ID)
	if err != nil {
		slog.Error("Failed to get loan account from db: ", err)
		return lam, err
	}

	currentState := lamdb.LoanState
	newState := lam.LoanState

	// check state is valid
	validState := rules.StateCheck(currentState, newState)
	if !validState {
		slog.Error("cannot go back to previous state")
		err = errors.New("cannot go back to previous state")
		return lam, err
	}

	if newState == "approved" {
		lam.ApproveDate = time.Now()
		lam.ValidateBy = "system"
	} else if newState == "invested" {
		var cm models.Customer
		cs := las.repo.GetCustomerRepository()
		cm, err = cs.GetCustomerById(ctx, lam.CustomerID)
		if err != nil {
			slog.Error("Cannot get customer")
			return lam, err
		}

		sender := utils.NewEmailSend()
		m := utils.NewEMailMessage("Amartha Agreement Docs", "Thank You for your participant in investing our borrower. Here is the agreement docs")
		m.To = []string{cm.Email}
		m.EmailAttachFile("../file/agreement/AgreementDoc.pdf")
		sender.SendEmail(m)
	} else if newState == "disbursed" {
		lam.DisbursementDate = time.Now()
		lam.DisbursementOfficer = "system"
	}

	lam.UpdateBy = "system"
	lam.UpdateDate = time.Now()

	lam, err = lar.Update(ctx, lam)
	if err != nil {
		slog.Error("Error update loan account: ", err)
		return lam, err
	}
	return lam, err
}

func (las *loanAccountService) UpdateProposedLoan(ctx context.Context, lam models.LoanAccount) (ln models.LoanAccount, err error) {
	lar := las.repo.GetLoanAccountRepository()

	lam.UpdateBy = "system"
	lam.UpdateDate = time.Now()

	ln, err = lar.UpdateProposedLoan(ctx, lam)
	if err != nil {
		slog.Error("Error update loan account: ", err)
		return ln, err
	}

	ln, err = lar.GetLoanAccountById(ctx, lam.ID)
	if err != nil {
		slog.Error("Error update loan account: ", err)
		return ln, err
	}

	return ln, err
}

func (las *loanAccountService) UploadDocument(ctx context.Context, doc models.Document, id uuid.UUID) (lam models.LoanAccount, err error) {
	lar := las.repo.GetLoanAccountRepository()

	docId := uuid.New()
	doc.ID = docId

	lam.ID = id
	lam.Documents = doc

	lam.UpdateBy = "system"
	lam.UpdateDate = time.Now()

	lam.ValidateDate = time.Now()
	lam.ValidateBy = "system"

	lam, err = lar.UpdateLoanDoc(ctx, lam)
	if err != nil {
		slog.Error("Error update loan account: ", err)
		return lam, err
	}

	lam, err = lar.GetLoanAccountById(ctx, id)
	if err != nil {
		slog.Error("Error update loan account: ", err)
		return lam, err
	}

	return lam, err
}
