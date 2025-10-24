package mydate

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

	return Date{
		month: month,
		year:  year,
	}, nil
}
