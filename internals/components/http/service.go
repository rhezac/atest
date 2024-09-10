package handlers

import (
	"github.com/labstack/echo/v4"

	"amartha-test/internals/components/customer"
	"amartha-test/internals/components/loan"

	_ "amartha-test/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

/*
	type svc struct {
		e    *echo.Echo
		addr string
	}

	func (s *svc) Start(ctx context.Context) {
		s.e.Start(s.addr)
	}

	func (s *svc) Stop(ctx context.Context) {
		s.e.Shutdown(ctx)
	}
*/

// @title Customer API
// @version 1.0
// @description This is a Customer API

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9000
// @BasePath /api/
func ApiRoutes(e *echo.Echo, cs customer.CustomerService, las loan.LoanAccountService) {
	api := e.Group("/api")
	{
		ch := customerApiHanlder{
			customerService: cs,
		}
		lah := loanAccountApiHanlder{
			loanAccountService: las,
		}

		customerApi := api.Group("/customer")
		customerApi.GET("/", ch.handlerFindAll)
		customerApi.POST("/", ch.handlerCreate)
		customerApi.PUT("/:id", ch.handlerUpdate)
		customerApi.GET("/:id", ch.handlerGetCustomerById)

		loanAccountApi := api.Group("/loan")
		loanAccountApi.GET("/", lah.handlerFindAll)
		loanAccountApi.POST("/", lah.handlerProposedLoan)
		loanAccountApi.PATCH("/:id", lah.handlerUpdateProposedLoan)
		loanAccountApi.GET("/:id", lah.handlerGetLoanAccountById)
		loanAccountApi.PUT("/:id", lah.handlerUpdate)
		loanAccountApi.POST("/:id/documents/", lah.handlerUploadDoc)
		loanAccountApi.POST("/:id/invest/", lah.handlerLoanInvest)

		//docLoanApi := loanAccountApi.Group("/:id/documents")
		//docLoanApi.POST("/", lah.handlerUploadDoc)
		//docAccountApi.GET("/:filename", lah.handlerDownloadDoc)

		downloadApi := api.Group("/download")
		//docAccountApi.POST("/:id", lah.handlerUploadDoc)
		downloadApi.GET("/:filename", lah.handlerDownloadDoc)

	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
