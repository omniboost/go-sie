package sie_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/omniboost/go-sie"
)

func TestAmount(t *testing.T) {
	tests := map[string]string{
		"0":          "0.00",
		"-0.001":     "0.00",
		"1234.5":     "1234.50",
		"-113320.54": "-113320.54",
		"0.005":      "0.01",
		"-0.005":     "-0.01",
	}
	for in, want := range tests {
		if got := sie.RequireAmountFromString(in).String(); got != want {
			t.Errorf("Amount(%s) = %q, want %q", in, got, want)
		}
	}
}

func TestAmountUnmarshalJSON(t *testing.T) {
	for _, in := range []string{`-178.57`, `"-178.57"`} {
		a := sie.Amount{}
		if err := json.Unmarshal([]byte(in), &a); err != nil {
			t.Fatalf("%s: %v", in, err)
		}
		if got := a.String(); got != "-178.57" {
			t.Errorf("%s: got %q", in, got)
		}
	}
}

func TestQuoted(t *testing.T) {
	if got, want := sie.Quoted("Bar \"Ö\"\tC:\\1\n"), `"Bar \"Ö\" C:\1 "`; got != want {
		t.Errorf("Quoted = %s, want %s", got, want)
	}
}

func TestField(t *testing.T) {
	tests := map[string]string{
		"OMNIBOOST":     "OMNIBOOST",
		"0123":          "0123",
		"":              `""`,
		"Visma Compact": `"Visma Compact"`,
		"a{b}":          `"a{b}"`,
	}
	for in, want := range tests {
		if got := sie.Field(in); got != want {
			t.Errorf("Field(%q) = %s, want %s", in, got, want)
		}
	}
}

func TestVoucherMustBalance(t *testing.T) {
	v := sie.Voucher{
		Date: sie.Date{time.Date(2026, 9, 29, 0, 0, 0, 0, time.UTC)},
		Transactions: sie.Transactions{
			{Account: 3010, Amount: sie.RequireAmountFromString("-100.10")},
			{Account: 1510, Amount: sie.RequireAmountFromString("100.00")},
		},
	}
	if _, err := v.MarshalSIE(); err == nil {
		t.Fatal("expected an error for an unbalanced voucher")
	}

	// amounts are compared as they are written, rounded to two decimals
	v.Transactions = sie.Transactions{
		{Account: 3010, Amount: sie.RequireAmountFromString("-0.1")},
		{Account: 3011, Amount: sie.RequireAmountFromString("-0.2")},
		{Account: 1510, Amount: sie.RequireAmountFromString("0.3")},
	}
	if _, err := v.MarshalSIE(); err != nil {
		t.Fatal(err)
	}
}

func TestDateUnmarshalJSON(t *testing.T) {
	for _, in := range []string{`"2026-08-28"`, `"2026-08-28T00:00:00+02:00"`, `"20260828"`} {
		d := sie.Date{}
		if err := json.Unmarshal([]byte(in), &d); err != nil {
			t.Fatalf("%s: %v", in, err)
		}
		if got := d.String(); got != "20260828" {
			t.Errorf("%s: got %q", in, got)
		}
	}

	d := sie.Date{}
	if err := json.Unmarshal([]byte(`""`), &d); err != nil || !d.IsZero() {
		t.Errorf("empty date: %v %v", d, err)
	}
}

func TestVoucherOptionalFields(t *testing.T) {
	v := sie.Voucher{
		Date: sie.Date{time.Date(2026, 9, 29, 0, 0, 0, 0, time.UTC)},
		Transactions: sie.Transactions{
			{Account: 3010, Amount: sie.RequireAmountFromString("-100"), Text: "Logi"},
			{Account: 2620, Amount: sie.RequireAmountFromString("100"), Date: sie.Date{time.Date(2026, 9, 28, 0, 0, 0, 0, time.UTC)}},
		},
	}

	ss, err := v.MarshalSIE()
	if err != nil {
		t.Fatal(err)
	}

	got := []string{}
	for _, s := range ss {
		got = append(got, strings.Join(s, "|"))
	}

	want := []string{
		`#VER|""|""|20260929`,
		`{`,
		`#TRANS|3010|{}|-100.00|""|"Logi"`,
		`#TRANS|2620|{}|100.00|20260928`,
		`}`,
	}

	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("got:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}
