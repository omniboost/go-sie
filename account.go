package sie

import "fmt"

type Accounts []Account

func (aa Accounts) MarshalSIE() ([][]string, error) {
	ss := [][]string{}

	for _, a := range aa {
		tmp, err := a.MarshalSIE()
		if err != nil {
			return ss, err
		}
		ss = append(ss, tmp...)
	}

	return ss, nil
}

// Account is a #KONTO item: every account used in the vouchers must be declared
type Account struct {
	// Must be numeric
	// Mandatory
	Number int
	// Mandatory
	Name string
}

func (a Account) MarshalSIE() ([][]string, error) {
	return [][]string{
		[]string{"#KONTO", fmt.Sprint(a.Number), Quoted(a.Name)},
	}, nil
}
