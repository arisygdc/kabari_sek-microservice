package http

import (
	"strings"
	"time"
)

type IRequest interface {
	Valid() error
}

func ValidateRequest(req IRequest) error {
	return req.Valid()
}

func Trim(str string) string {
	return strings.Trim(str, " ")
}

func lengRule(min int, max int, str string) bool {
	return len(str) > max || len(str) < min
}

func IsAlphaNumeric(str string) bool {
	for _, v := range str {
		// Allowed lower alfabeth and numeric
		if (v < 48 || (v > 57 && v < 97)) || v > 122 {
			return false
		}
	}
	return true
}

func IsNumeric(str string) bool {
	for _, v := range str {
		if v < 48 || v > 57 {
			return false
		}
	}
	return true
}

// date format yyy-mm-dd
const dateLayout = "2006-02-02"

// beware ZERO before NUMBER
// that will reset its date
// input example "1988-7-22"
func ParseDate(str string) (time.Time, error) {
	return time.Parse(dateLayout, str)
}
