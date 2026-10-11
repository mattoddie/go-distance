// Package distance provides a Distance type for representing lengths, with
// conversions between metric and imperial units and a parser for strings such
// as "5km", "300ft" or "1mi200yd".
//
// A Distance is stored as a float64 number of metres:
//
//	d, err := distance.Parse("1mi200yd")
//	if err != nil {
//		// handle error
//	}
//	fmt.Println(d.Kilometres())
//
// Distance implements [encoding/json.Marshaler] and
// [encoding/json.Unmarshaler], encoding as a JSON string such as "5000m".
package distance

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"
)

// Distance represents a length as a float64 number of metres.
type Distance float64

// Common distances. To count the number of units in a Distance, divide:
//
//	ft := distance.Foot
//	fmt.Print(float64(ft / distance.Inch)) // prints 12
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

// String returns the distance in metres, formatted as "<n>m", for example
// "1609.344m". It never uses exponent notation, so the result can always be
// read back by [Parse] unless the distance is NaN or infinite.
func (d Distance) String() string {
	return strconv.FormatFloat(float64(d), 'f', -1, 64) + "m"
}

// Metres returns the distance as a floating point number of metres.
func (d Distance) Metres() float64 {
	return float64(d)
}

// Centimetres returns the distance as a floating point number of centimetres.
func (d Distance) Centimetres() float64 {
	return float64(d / Centimetre)
}

// Millimetres returns the distance as a floating point number of millimetres.
func (d Distance) Millimetres() float64 {
	return float64(d / Millimetre)
}

// Kilometres returns the distance as a floating point number of kilometres.
func (d Distance) Kilometres() float64 {
	return float64(d / Kilometre)
}

// Miles returns the distance as a floating point number of miles.
func (d Distance) Miles() float64 {
	return float64(d / Mile)
}

// Inches returns the distance as a floating point number of inches.
func (d Distance) Inches() float64 {
	return float64(d / Inch)
}

// Yards returns the distance as a floating point number of yards.
func (d Distance) Yards() float64 {
	return float64(d / Yard)
}

// Feet returns the distance as a floating point number of feet.
func (d Distance) Feet() float64 {
	return float64(d / Foot)
}

// UnmarshalJSON implements [encoding/json.Unmarshaler]. It accepts a JSON
// string in any format understood by [Parse]. A JSON null leaves the
// distance unchanged.
func (d *Distance) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	v, err := Parse(s)
	if err != nil {
		return err
	}
	*d = v
	return nil
}

// MarshalJSON implements [encoding/json.Marshaler]. The distance is encoded
// as a JSON string in the format returned by [Distance.String]. It returns
// an error if the distance is NaN or infinite.
func (d Distance) MarshalJSON() ([]byte, error) {
	f := float64(d)
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return nil, errors.New("distance: cannot marshal " + d.String())
	}
	return []byte(strconv.Quote(d.String())), nil
}
