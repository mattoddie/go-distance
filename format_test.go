package distance

import (
	"testing"
)

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
			t.Errorf("Wrong distance, got %v expected %v", calculated, tc.expected)
		}
	}
}
