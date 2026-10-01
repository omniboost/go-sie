package sie

import "fmt"

type Dimensions []Dimension

func (dd Dimensions) MarshalSIE() ([][]string, error) {
	ss := [][]string{}

	for _, d := range dd {
		tmp, err := d.MarshalSIE()
		if err != nil {
			return ss, err
		}
		ss = append(ss, tmp...)
	}

	return ss, nil
}

// Dimension is a #DIM item, 1 is reserved for cost centres (resultatenhet)
type Dimension struct {
	// Mandatory
	Number int
	// Mandatory
	Name string
}

func (d Dimension) MarshalSIE() ([][]string, error) {
	return [][]string{
		[]string{"#DIM", fmt.Sprint(d.Number), Quoted(d.Name)},
	}, nil
}

type Objects []Object

func (oo Objects) MarshalSIE() ([][]string, error) {
	ss := [][]string{}

	for _, o := range oo {
		tmp, err := o.MarshalSIE()
		if err != nil {
			return ss, err
		}
		ss = append(ss, tmp...)
	}

	return ss, nil
}

// Object is an #OBJEKT item: a value within a dimension
type Object struct {
	// Mandatory
	Dimension int
	// Mandatory
	Number string
	// Mandatory
	Name string
}

func (o Object) MarshalSIE() ([][]string, error) {
	return [][]string{
		[]string{"#OBJEKT", fmt.Sprint(o.Dimension), Field(o.Number), Quoted(o.Name)},
	}, nil
}
