package sie

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// Date is written as YYYYMMDD
type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return json.Marshal("")
	}
	return json.Marshal(d.Format("2006-01-02"))
}

func (d Date) String() string {
	if d.IsZero() {
		return ""
	}
	return d.Format("20060102")
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try iso8601 date format
	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return nil
	}

	// try sie date format
	d.Time, err = time.Parse("20060102", value)
	return err
}

// Amount is written with a decimal point and two decimals, debit positive and
// credit negative. It unmarshals from a JSON number or a quoted string.
type Amount struct {
	decimal.Decimal
}

func NewAmount(d decimal.Decimal) Amount {
	return Amount{d}
}

// NewAmountFromString parses a decimal string such as "-113320.54"
func NewAmountFromString(s string) (Amount, error) {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return Amount{}, err
	}
	return Amount{d}, nil
}

// RequireAmountFromString parses a decimal string and panics when it is not a
// valid number. Use it for constants, not for input.
func RequireAmountFromString(s string) Amount {
	return Amount{decimal.RequireFromString(s)}
}

// Rounded returns the amount rounded to two decimals, as it is written
func (a Amount) Rounded() decimal.Decimal {
	return a.Decimal.Round(2)
}

func (a Amount) String() string {
	r := a.Rounded()
	if r.IsZero() {
		return "0.00"
	}
	return r.StringFixed(2)
}

// Quoted writes a string between double quotes. Embedded quotes are preceded by
// a backslash and control characters (ASCII 0-31 and 127) are replaced by a
// space, see 5.7 of the spec.
func Quoted(s string) string {
	s = strings.Map(func(r rune) rune {
		if r < 32 || r == 127 {
			return ' '
		}
		return r
	}, s)
	s = strings.ReplaceAll(s, `"`, `\"`)
	return `"` + s + `"`
}

// Field writes a string as is, unless it is empty or contains characters that
// would break the item (spaces, quotes, braces, control characters): then it is
// quoted.
func Field(s string) string {
	if s == "" || strings.IndexFunc(s, func(r rune) bool {
		return r <= 32 || r == 127 || r == '"' || r == '{' || r == '}'
	}) >= 0 {
		return Quoted(s)
	}
	return s
}
