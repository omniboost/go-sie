package sie

import (
	"bytes"
	"strings"
)

// File is a SIE 4I file. Items are written in the order the format requires.
type File struct {
	Header     Header
	Accounts   Accounts
	Dimensions Dimensions
	Objects    Objects
	Vouchers   Vouchers
}

type Section interface {
	MarshalSIE() ([][]string, error)
}

func (f File) MarshalSIE() ([][]string, error) {
	ss := [][]string{}

	for _, s := range []Section{f.Header, f.Accounts, f.Dimensions, f.Objects, f.Vouchers} {
		tmp, err := s.MarshalSIE()
		if err != nil {
			return ss, err
		}
		ss = append(ss, tmp...)
	}

	return ss, nil
}

// Marshal writes the file as tab separated fields with CRLF line endings.
// Converting to the target encoding (e.g. windows-1252) is up to the caller.
func Marshal(f File) ([]byte, error) {
	ss, err := f.MarshalSIE()
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	for _, s := range ss {
		b.WriteString(strings.Join(s, "\t"))
		b.WriteString("\r\n")
	}

	return b.Bytes(), nil
}
