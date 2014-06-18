package distance

import (
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
		if out := tc.in.String(); out != tc.expected {
			t.Errorf("Wrong string, got %v expected %v", out, tc.in)
		}
	}
}

func TestParseDistanceToMetres(t *testing.T) {
	testCases := []struct {
		in       string
		expected float64
	}{
		{"5m", 5},
		{"-5m", -5},
		{"10m", 10},
		{"120cm", 1.2},
		{"1245mm", 1.245},
		{"2mi", 3218.688},
		{"2km", 2000},
		{"100yd", 91.44},
		{"300ft", 91.44},
		{"3600in", 91.44000000000001},
	}

	for _, tc := range testCases {
		calculated, err := ParseDistance(tc.in)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if calculated.Metres() != tc.expected {
			t.Errorf("Wrong distance, got %v expected %v", calculated.Metres(), tc.expected)
		}
	}
}

func TestParseDistanceToMiles(t *testing.T) {
	testCases := []struct {
		in       string
		expected float64
	}{
		{"5mi", 5},
		{"5m", 0.0031068559611866697},
		{"1760yd", 1},
	}

	for _, tc := range testCases {
		calculated, err := ParseDistance(tc.in)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if calculated.Miles() != tc.expected {
			t.Errorf("Wrong distance, got %v expected %v", calculated.Miles(), tc.expected)
		}
	}
}

func TestBadParses(t *testing.T) {
	testCases := []struct {
		in string
	}{
		{"5"},
		{"5m5"},
	}

	for _, tc := range testCases {
		_, err := ParseDistance(tc.in)
		if err == nil {
			t.Errorf("Expected error for %v", tc.in)
		}
	}
}
