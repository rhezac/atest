package main

import (
	cservice "amartha-test/internals/components/customer/service"
	handlers "amartha-test/internals/components/http"
	laservice "amartha-test/internals/components/loan/service"
	"amartha-test/internals/components/repository"
	"amartha-test/internals/utils"
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"

	_ "amartha-test/docs"
)

func main() {
	// Set config
	config, err := utils.LoadConfig("../config")
	if err != nil {
		fmt.Println("Could not load environment variables", err)
	}

	// Set log
	utils.InitLog(&config)
	slog.Debug("Log Set")

	// Set db
	db, err := utils.ConnectDatabase(&config)
	if err != nil {
		slog.Error("Could not connect to database", err)
	}

	// Set repository
	repo := repository.NewRepository(db)

	// Set service
	cs := cservice.NewCustomerService(repo)
	las := laservice.NewLoanAccountService(repo)

	// Set Server
	e := echo.New()
	handlers.ApiRoutes(e, cs, las)
	e.Start(":9000")

	slog.Debug("App Started")

}
