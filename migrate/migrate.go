package main

import (
	"amartha-test/internals/utils"
	"fmt"

	"golang.org/x/exp/slog"
)

func main() {
	// Set config
	config, err := utils.LoadConfig("../config")
	if err != nil {
		fmt.Println("Could not load environment variables", err)
	}

	// Set db
	db, err := utils.ConnectDatabase(&config)
	if err != nil {
		slog.Error("Could not connect to database", err)
	}

	utils.PrepareDatabase(db)
	slog.Info("Done Migrate")
}
