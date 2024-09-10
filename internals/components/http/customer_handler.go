package handlers

import (
	"amartha-test/internals/components/customer"
	"amartha-test/internals/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
)

type customerApiHanlder struct {
	customerService customer.CustomerService
}

type CustomerResponseID struct {
	ID string `json:"id"`
}

// @Tags Customer
// @Summary Create Csutomer
// @Description Create Customer
// @Router /customer [post]
// @Param request body id true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Success 200 {array} Customer
// @Failure 400
func (ch *customerApiHanlder) handlerCreate(c echo.Context) error {
	var customer models.Customer
	var customerId CustomerResponseID

	ctx := c.Request().Context()

	err := c.Bind(&customer)
	if err != nil {
		slog.Error("Failed to bind customer data: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	customerID, err := ch.customerService.Create(ctx, customer)
	if err != nil {
		slog.Error("Failed to create customer: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	customerId.ID = customerID

	return c.JSON(http.StatusOK, customerId)
}

func (ch *customerApiHanlder) handlerFindAll(c echo.Context) error {
	var customer []models.Customer

	ctx := c.Request().Context()

	customer, err := ch.customerService.FindAll(ctx)
	if err != nil {
		slog.Error("Failed to get customers data: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, customer)
}

func (ch *customerApiHanlder) handlerUpdate(c echo.Context) error {
	var customer models.Customer

	ctx := c.Request().Context()

	err := c.Bind(&customer)
	if err != nil {
		slog.Error("Failed to bind customer data: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	id := uuid.MustParse(c.Param("id"))
	customer.ID = id
	customer, err = ch.customerService.Update(ctx, customer)
	if err != nil {
		slog.Error("Failed to update customer: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, customer)
}

func (ch *customerApiHanlder) handlerGetCustomerById(c echo.Context) error {
	var customer models.Customer

	ctx := c.Request().Context()

	id := uuid.MustParse(c.Param("id"))
	customer, err := ch.customerService.GetCustomerById(ctx, id)
	if err != nil {
		slog.Error("Failed to get customer: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, customer)
}
