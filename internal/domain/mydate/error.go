package mydate

import "fmt"

var ErrFutureDate = fmt.Errorf("date is in the future")

type ErrInvalidMonth struct {
	Month int
}

func NewErrInvalidMonth(month int) error {
	return ErrInvalidMonth{
		Month: month,
	}
}

func (e ErrInvalidMonth) Error() string {
	return fmt.Sprintf("invalid month: %d", e.Month)
}
