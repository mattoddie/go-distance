package distance

import (
	"math"
	"testing"
)

func FuzzParse(f *testing.F) {
	for _, s := range []string{
		"0", "5m", "-5m", "+1.5km", ".5m", "5.m", "1mi200yd", "3600in",
		"0.0000001m", "1000000000000000000000m", "", "-", ".", "5", "5m5", "1e5m",
	} {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		d, err := Parse(s)
		if err != nil {
			return
		}

		f := float64(d)
		if math.IsNaN(f) || math.IsInf(f, 0) {
			t.Fatalf("Parse(%q) returned %v", s, f)
		}

		// Anything Parse accepts must round trip through String.
		again, err := Parse(d.String())
		if err != nil {
			t.Fatalf("Parse(%q) = %v, but Parse(%q) failed: %v", s, d, d.String(), err)
		}
		if again != d {
			t.Fatalf("Parse(%q) = %v, but Parse(String()) = %v", s, d, again)
		}
	})
}
