// Portions of this file are derived from the Go standard library's time
// package. Copyright 2010 The Go Authors. Use of that source code is governed
// by a BSD-style license that can be found in the LICENSE-GO file.

package distance

import (
	"errors"
	"math"
	"strconv"
)

var unitMap = map[string]Distance{
	"m":  Metre,
	"cm": Centimetre,
	"mi": Mile,
	"mm": Millimetre,
	"km": Kilometre,
	"in": Inch,
	"ft": Foot,
	"yd": Yard,
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// leadingDigits returns the number of leading [0-9] bytes in s.
func leadingDigits(s string) int {
	i := 0
	for i < len(s) && isDigit(s[i]) {
		i++
	}
	return i
}

// Parse parses a distance string. A distance string is a possibly signed
// sequence of decimal numbers, each with optional fraction and a unit suffix,
// such as "300ft", "-1.5km" or "1mi200yd". Valid units are "mm", "cm", "m",
// "km", "in", "ft", "yd" and "mi".
//
// Parse accepts any string produced by [Distance.String].
func Parse(s string) (Distance, error) {
	// [-+]?([0-9]*(\.[0-9]*)?[a-z]+)+
	orig := s
	var d Distance
	neg := false

	// Consume [-+]?
	if s != "" {
		c := s[0]
		if c == '-' || c == '+' {
			neg = c == '-'
			s = s[1:]
		}
	}
	// Special case: if all that is left is "0", this is zero.
	if s == "0" {
		return 0, nil
	}
	if s == "" {
		return 0, errors.New("distance: invalid distance " + strconv.Quote(orig))
	}
	for s != "" {
		// The next character must be [0-9.]
		if s[0] != '.' && !isDigit(s[0]) {
			return 0, errors.New("distance: invalid distance " + strconv.Quote(orig))
		}

		// Consume [0-9]*(\.[0-9]*)?
		n := leadingDigits(s)
		digits := n
		if n < len(s) && s[n] == '.' {
			frac := leadingDigits(s[n+1:])
			digits += frac
			n += 1 + frac
		}
		if digits == 0 {
			// no digits (e.g. ".m" or "-.m")
			return 0, errors.New("distance: invalid distance " + strconv.Quote(orig))
		}
		g, err := strconv.ParseFloat(s[:n], 64)
		if err != nil {
			return 0, errors.New("distance: invalid distance " + strconv.Quote(orig))
		}
		s = s[n:]

		// Consume unit.
		i := 0
		for ; i < len(s); i++ {
			c := s[i]
			if c == '.' || isDigit(c) {
				break
			}
		}
		if i == 0 {
			return 0, errors.New("distance: missing unit in distance " + strconv.Quote(orig))
		}
		u := s[:i]
		s = s[i:]
		unit, ok := unitMap[u]
		if !ok {
			return 0, errors.New("distance: unknown unit " + strconv.Quote(u) + " in distance " + strconv.Quote(orig))
		}

		d += Distance(g) * unit
	}

	if math.IsInf(float64(d), 0) {
		return 0, errors.New("distance: invalid distance " + strconv.Quote(orig))
	}
	if neg {
		d = -d
	}
	return d, nil
}
