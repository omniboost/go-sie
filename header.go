package sie

import "fmt"

// Header holds the identification items at the top of the file
type Header struct {
	// #FLAGGA, 0 means the file has not been imported yet
	// Mandatory
	Flagga int
	// #PROGRAM, name and version of the program that created the file
	// Mandatory
	Program        string
	ProgramVersion string
	// #FORMAT, character set: PC8
	// Mandatory
	Format string
	// #GEN, date the file was created
	// Mandatory
	Gen Date
	// #SIETYP, 4 for SIE 4I (vouchers only)
	// Mandatory
	SieTyp int
	// #PROSA, free text
	Prosa string
	// #FNAMN, company name
	// Mandatory
	FNamn string
	// #KPTYP, type of chart of accounts: BAS95, BAS96, EUBAS97, NE2007
	KPTyp string
}

func (h Header) MarshalSIE() ([][]string, error) {
	ss := [][]string{}

	ss = append(ss, []string{"#FLAGGA", fmt.Sprint(h.Flagga)})

	if h.ProgramVersion != "" {
		ss = append(ss, []string{"#PROGRAM", Field(h.Program), Field(h.ProgramVersion)})
	} else {
		ss = append(ss, []string{"#PROGRAM", Field(h.Program)})
	}

	ss = append(ss, []string{"#FORMAT", h.Format})
	ss = append(ss, []string{"#GEN", h.Gen.String()})
	ss = append(ss, []string{"#SIETYP", fmt.Sprint(h.SieTyp)})

	if h.Prosa != "" {
		ss = append(ss, []string{"#PROSA", Quoted(h.Prosa)})
	}

	ss = append(ss, []string{"#FNAMN", Quoted(h.FNamn)})

	if h.KPTyp != "" {
		ss = append(ss, []string{"#KPTYP", h.KPTyp})
	}

	return ss, nil
}
