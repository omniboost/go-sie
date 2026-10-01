package sie

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

type Vouchers []Voucher

func (vv Vouchers) MarshalSIE() ([][]string, error) {
	ss := [][]string{}

	for _, v := range vv {
		tmp, err := v.MarshalSIE()
		if err != nil {
			return ss, err
		}
		ss = append(ss, tmp...)
	}

	return ss, nil
}

// Voucher is a #VER item followed by its #TRANS rows between braces. The
// transactions of a voucher must balance (sum to zero).
type Voucher struct {
	// Empty means the receiving system picks the series
	Series string
	// Empty means the receiving system numbers the voucher
	Number string
	// Mandatory
	Date Date
	Text string
	Transactions
}

func (v Voucher) MarshalSIE() ([][]string, error) {
	ss := [][]string{}

	if total := v.Transactions.Total(); !total.IsZero() {
		return ss, fmt.Errorf("voucher %s %s %s does not balance: off by %s", v.Series, v.Number, v.Date, total.StringFixed(2))
	}

	row := []string{"#VER", Quoted(v.Series), Field(v.Number), v.Date.String()}
	if v.Text != "" {
		row = append(row, Quoted(v.Text))
	}
	ss = append(ss, row)

	ss = append(ss, []string{"{"})

	tmp, err := v.Transactions.MarshalSIE()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	ss = append(ss, []string{"}"})

	return ss, nil
}

type Transactions []Transaction

func (tt Transactions) MarshalSIE() ([][]string, error) {
	ss := [][]string{}

	for _, t := range tt {
		tmp, err := t.MarshalSIE()
		if err != nil {
			return ss, err
		}
		ss = append(ss, tmp...)
	}

	return ss, nil
}

// Total returns the sum of the transaction amounts, each amount rounded to two
// decimals as it is written
func (tt Transactions) Total() decimal.Decimal {
	total := decimal.Zero
	for _, t := range tt {
		total = total.Add(t.Amount.Rounded())
	}
	return total
}

// Transaction is a #TRANS row
type Transaction struct {
	// Must be numeric
	// Mandatory
	Account int
	Objects TransactionObjects
	// Mandatory
	Amount Amount
	// Optional, only written when set
	Date Date
	Text string
}

func (t Transaction) MarshalSIE() ([][]string, error) {
	row := []string{"#TRANS", fmt.Sprint(t.Account), t.Objects.String(), t.Amount.String()}

	if !t.Date.IsZero() || t.Text != "" {
		row = append(row, t.Date.String())
		if row[len(row)-1] == "" {
			row[len(row)-1] = Quoted("")
		}
	}
	if t.Text != "" {
		row = append(row, Quoted(t.Text))
	}

	return [][]string{row}, nil
}

// TransactionObjects is the object list of a #TRANS row: {"1"  "307"}
type TransactionObjects []TransactionObject

func (oo TransactionObjects) String() string {
	pairs := make([]string, len(oo))
	for i, o := range oo {
		pairs[i] = Quoted(fmt.Sprint(o.Dimension)) + "  " + Quoted(o.Number)
	}
	return "{" + strings.Join(pairs, "  ") + "}"
}

type TransactionObject struct {
	Dimension int
	Number    string
}
