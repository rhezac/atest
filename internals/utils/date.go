package utils

import (
	"time"

	"golang.org/x/exp/slog"
)

func ValidateDateFormat(dates string) (dateValid time.Time, err error) {
	dateValid, err = time.Parse("01/02/2006", dates) // the format is mm/dd/yyyy
	if err != nil {
		slog.Error("The format is not valid", err)
	}

	return
}
