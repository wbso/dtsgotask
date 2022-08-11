package app

import (
	"errors"
	"fmt"
	"time"
)

var ErrInvalidInput = errors.New("invalid input")

func InvalidInputError(op string) error {
	return fmt.Errorf("OperationNotPermit %w : %s", ErrInvalidInput, op)
}

func parseYMDToTime(ymd string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", ymd)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: %w", err)
	}
	return t, nil
}
