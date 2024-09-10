package handlers

import (
	"amartha-test/internals/components/loan"
	"amartha-test/internals/models"
	"amartha-test/internals/utils"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"golang.org/x/exp/slog"
)

type loanAccountApiHanlder struct {
	loanAccountService loan.LoanAccountService
}

type GetLoanAccountRequest struct {
	ID                  uuid.UUID       `json:"id"`
	CustomerID          uuid.UUID       `json:"customer_id"`
	LoanProductID       uuid.UUID       `json:"loan_product_id"`
	Rate                decimal.Decimal `json:"rate"`
	ValidateDate        string          `json:"validate_date"`
	ValidateBy          string          `json:"validate_by"`
	ApproveDate         string          `json:"approve_date"`
	AgreementNo         string          `json:"agreement_no"`
	AgreementDate       string          `json:"agreement_date"`
	AgreementDocID      uuid.UUID       `json:"agreement_doc_id"`
	DisbursementDate    string          `json:"disbursement_date"`
	DisbursementOfficer string          `json:"disbursement_officer"`
	Principal           decimal.Decimal `json:"principal_amount"`
	TotalAmount         decimal.Decimal `json:"total_amount"`
	OutstandingAmount   decimal.Decimal `json:"outstanding_amount"`
	ROI                 decimal.Decimal `json:"roi"`
	TenorValue          int             `json:"tenor_value"`
	Status              string          `json:"loan_status"`
	LoanState           string          `json:"loan_state"`
	Kind                string          `json:"kind" gorm:"json"`
	Investors           string          `json:"investors" gorm:"json"`
	CreateBy            string          `json:"-" gorm:"varchar(50)"`
	UpdateBy            string          `json:"-" gorm:"varchar(50)"`
	Documents           string          `json:"documents" gorm:"json"`
}

type LoanResponseID struct {
	ID string `json:"id"`
}

func (lah *loanAccountApiHanlder) handlerProposedLoan(c echo.Context) error {
	var loanAccount models.LoanAccount
	var loanIdResponse LoanResponseID

	ctx := c.Request().Context()

	err := c.Bind(&loanAccount)
	if err != nil {
		slog.Error("Failed to bind loan account data: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	loanAccountID, err := lah.loanAccountService.Create(ctx, loanAccount)
	if err != nil {
		slog.Error("Failed to create loan account: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	loanIdResponse.ID = loanAccountID

	return c.JSON(http.StatusOK, loanIdResponse)
}

func (lah *loanAccountApiHanlder) handlerFindAll(c echo.Context) error {
	var loanAccounts []models.LoanAccount

	ctx := c.Request().Context()

	loanAccounts, err := lah.loanAccountService.FindAll(ctx)
	if err != nil {
		slog.Error("Failed to get loan account data: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, loanAccounts)
}

func (lah *loanAccountApiHanlder) handlerUpdate(c echo.Context) error {
	var loanAccount models.LoanAccount

	ctx := c.Request().Context()

	err := c.Bind(&loanAccount)
	if err != nil {
		slog.Error("Failed to bind loan account data: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	loanAccount.ID = uuid.MustParse(c.Param("id"))

	loanAccounts, err := lah.loanAccountService.Update(ctx, loanAccount)
	if err != nil {
		slog.Error("Failed to update loan account: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, loanAccounts)
}

func (lah *loanAccountApiHanlder) handlerUpdateProposedLoan(c echo.Context) error {
	var loanAccount models.LoanAccount

	ctx := c.Request().Context()

	err := c.Bind(&loanAccount)
	if err != nil {
		slog.Error("Failed to bind loan account data: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	loanAccount.ID = uuid.MustParse(c.Param("id"))

	loanAccounts, err := lah.loanAccountService.UpdateProposedLoan(ctx, loanAccount)
	if err != nil {
		slog.Error("Failed to update loan account: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, loanAccounts)
}

func (lah *loanAccountApiHanlder) handlerGetLoanAccountById(c echo.Context) error {
	var loanAccount models.LoanAccount

	ctx := c.Request().Context()

	id := uuid.MustParse(c.Param("id"))
	loanAccount, err := lah.loanAccountService.GetLoanAccountById(ctx, id)
	if err != nil {
		slog.Error("Failed to get loan account: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, loanAccount)
}

func (lah *loanAccountApiHanlder) handlerUploadDoc(c echo.Context) error {
	var document models.Document
	var loanAccount models.LoanAccount

	ctx := c.Request().Context()

	err := c.Bind(&document)
	if err != nil {
		slog.Error("Failed to bind customer data: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	loanId := uuid.MustParse(c.Param("id"))

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	fileName := file.Filename
	document.Name = fileName
	filePath := "../file/document/" + fileName

	err = utils.UploadDoc(file, filePath)
	if err != nil {
		slog.Error("Failed to upload file: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	document.DocUrl = "http://localhost:9000/api/download/" + fileName

	desc := c.FormValue("description")
	document.Description = desc

	docDate := c.FormValue("document_date")
	document.DocumentDate, err = utils.ValidateDateFormat(docDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	uploadBy := c.FormValue("am_name")
	document.UploadBy = uploadBy

	loanAccount, err = lah.loanAccountService.UploadDocument(ctx, document, loanId)
	if err != nil {
		slog.Error("Failed to upload document: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, loanAccount)
}

func (lah *loanAccountApiHanlder) handlerDownloadDoc(c echo.Context) error {
	var document models.Document

	fileName := c.Param("filename")

	filePath := "../file/document/" + fileName

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return echo.NewHTTPError(http.StatusNotFound, "File not found")
	}

	c.File(filePath)

	return c.JSON(http.StatusOK, document)
}

func (lah *loanAccountApiHanlder) handlerLoanInvest(c echo.Context) error {
	var loanAccount models.LoanAccount
	var loanInvest models.LoanInvested

	ctx := c.Request().Context()

	err := c.Bind(&loanInvest)
	if err != nil {
		slog.Error("Failed to bind loan invest data: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	id := uuid.MustParse(c.Param("id"))

	loanAccount, err = lah.loanAccountService.LoanInvest(ctx, loanInvest, id)
	if err != nil {
		slog.Error("Failed to update loan account: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, loanAccount)
}
