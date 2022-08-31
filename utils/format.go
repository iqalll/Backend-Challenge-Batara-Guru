package utils

import (
	"strings"

	humanize "github.com/dustin/go-humanize"
)

func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp " + stringValue + ",00"
}

func FormatDateString(date string) string {
	date = strings.Split(date, "T")[0]
	return date
}
