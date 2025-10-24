package mydate

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	myDateExactRegex = regexp.MustCompile(`^(\d{2})\.(\d{4})$`)
	myDateRegex      = regexp.MustCompile(`(\d{1,2})\.(\d{4})`)
)

type Date struct {
	month   int
	year    int
	current bool
}

func New(month, year int) (Date, error) {
	if month < 1 || month > 12 {
		return Date{}, fmt.Errorf("invalid month: %d", month)
	}

	if inFuture(month, year) {
		return Date{}, ErrFutureDate
	}

	return Date{
		month: month,
		year:  year,
	}, nil
}

func Current() Date {
	return Date{
		current: true,
	}
}

func (d *Date) Equal(other Date) bool {
	if d.current && other.current {
		return true
	}

	if d.current || other.current {
		return false
	}

	return d.month == other.month && d.year == other.year
}

func (d *Date) Less(other Date) bool {
	if d.current && other.current {
		return false
	}

	if d.current {
		return false
	}

	if other.current {
		return true
	}

	if d.year != other.year {
		return d.year < other.year
	}

	return d.month < other.month
}

func (d *Date) Since(other Date) int {
	if d.current && other.current {
		return 0
	}

	if d.Less(other) {
		return -other.Since(*d)
	}

	y, m := d.year, d.month
	if d.current {
		now := time.Now()
		y = now.Year()
		m = int(now.Month())
	}

	years := y - other.year
	months := m - other.month

	return years*12 + months
}

func (d *Date) Current() bool {
	return d.current
}

func (d *Date) String() string {
	if d.current {
		return "настоящее время"
	}

	return fmt.Sprintf("%02d.%d", d.month, d.year)
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	date, err := ParseDate(str)
	if err != nil {
		return err
	}

	d.current = date.current
	d.month = date.month
	d.year = date.year

	return nil
}

func ParseDate(str string) (Date, error) {
	if str == "настоящее время" {
		return Date{
			current: true,
		}, nil
	}

	matches := myDateExactRegex.FindStringSubmatch(str)
	if len(matches) != 3 {
		return Date{}, fmt.Errorf("invalid date format: %s", str)
	}

	month, err := strconv.Atoi(matches[1])
	if err != nil || month < 1 || month > 12 {
		return Date{}, fmt.Errorf("invalid month in date: %s", str)
	}

	year, err := strconv.Atoi(matches[2])
	if err != nil {
		return Date{}, fmt.Errorf("invalid year in date: %s", str)
	}

	return Date{
		month: month,
		year:  year,
	}, nil
}

func ExtractAndParseDate(str string) (Date, error) {
	if strings.Contains(str, "настоящее время") {
		return Date{
			current: true,
		}, nil
	}

	matches := myDateRegex.FindStringSubmatch(str)
	if len(matches) != 3 {
		return Date{}, fmt.Errorf("invalid date format: %s", str)
	}

	month, err := strconv.Atoi(matches[1])
	if err != nil || month < 1 || month > 12 {
		return Date{}, fmt.Errorf("invalid month in date: %s", str)
	}

	year, err := strconv.Atoi(matches[2])
	if err != nil {
		return Date{}, fmt.Errorf("invalid year in date: %s", str)
	}

	if inFuture(month, year) {
		return Date{}, ErrFutureDate
	}

	return Date{
		month: month,
		year:  year,
	}, nil
}

func inFuture(month, year int) bool {
	now := time.Now()
	return year > now.Year() || (year == now.Year() && month > int(now.Month()))
}
