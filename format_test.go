package distance

import (
	"strings"
	"testing"
)

func TestToString(t *testing.T) {
	testCases := []struct {
		in       Distance
		expected string
	}{
		{5 * Metre, "5m"},
		{2 * Mile, "3218.688m"},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			if out := tc.in.String(); out != tc.expected {
				t.Errorf("Wrong string, got %v expected %v", out, tc.expected)
			}
		})
	}
}

func TestParseToMetres(t *testing.T) {
	testCases := []struct {
		in       string
		expected float64
	}{
		{"5m", 5},
		{"-5m", -5},
		{"+5m", 5},
		{"0", 0},
		{"10m", 10},
		{"120cm", 1.2},
		{"1245mm", 1.245},
		{"2mi", 3218.688},
		{"2km", 2000},
		{"1.5km", 1500},
		{".5m", 0.5},
		{"1km500m", 1500},
		{"100yd", 91.44},
		{"300ft", 91.44},
		{"3600in", 91.44000000000001},
	}

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			calculated, err := Parse(tc.in)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if calculated.Metres() != tc.expected {
				t.Errorf("Wrong distance, got %v expected %v", calculated.Metres(), tc.expected)
			}
		})
	}
}

func TestParseToMiles(t *testing.T) {
	testCases := []struct {
		in       string
		expected float64
	}{
		{"5mi", 5},
		{"5m", 0.0031068559611866697},
		{"1760yd", 1},
	}

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			calculated, err := Parse(tc.in)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if calculated.Miles() != tc.expected {
				t.Errorf("Wrong distance, got %v expected %v", calculated.Miles(), tc.expected)
			}
		})
	}
}

func TestBadParses(t *testing.T) {
	testCases := []string{
		"",
		"-",
		"5",
		"5m5",
		".m",
		"5furlong",
		"m",
	}

	for _, in := range testCases {
		t.Run(in, func(t *testing.T) {
			if _, err := Parse(in); err == nil {
				t.Errorf("Expected error for %q", in)
			}
		})
	}
}

func TestParseLongNumbers(t *testing.T) {
	testCases := []struct {
		in       string
		expected Distance
	}{
		{"0.0000001m", 1e-7},
		{"1000000000000000000000m", 1e21},
		{"0.12345678901234567890123m", 0.12345678901234567890123},
	}

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			got, err := Parse(tc.in)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if got != tc.expected {
				t.Errorf("Wrong distance, got %v expected %v", got, tc.expected)
			}
		})
	}
}

func TestParseOverflow(t *testing.T) {
	in := "1" + strings.Repeat("0", 400) + "m"
	if _, err := Parse(in); err == nil {
		t.Error("Expected error for overflowing distance")
	}
}
