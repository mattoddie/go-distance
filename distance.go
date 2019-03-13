package distance

import (
	"fmt"
	"strings"
)

type Distance float64

const (
	Metre      Distance = 1
	Centimetre          = Metre / 100
	Millimetre          = Centimetre / 10
	Kilometre           = 1000 * Metre
	Mile                = 1.609344 * Kilometre
	Yard                = Mile / 1760
	Foot                = Yard / 3
	Inch                = Foot / 12
)

func (d Distance) String() string {
	return fmt.Sprintf("%vm", float64(d))
}

func (d Distance) Metres() float64 {
	return float64(d)
}

func (d Distance) Centimetres() float64 {
	return float64(d / Centimetre)
}

func (d Distance) Millimetres() float64 {
	return float64(d / Millimetre)
}

func (d Distance) Kilometres() float64 {
	return float64(d / Kilometre)
}

func (d Distance) Miles() float64 {
	return float64(d / Mile)
}

func (d Distance) Inches() float64 {
	return float64(d / Inch)
}

func (d Distance) Yards() float64 {
	return float64(d / Yard)
}

func (d Distance) Feet() float64 {
	return float64(d / Foot)
}

func (d *Distance) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if s == "null" {
		return
	}

	*d, err = ParseDistance(s)
	return
}

func (d *Distance) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}
