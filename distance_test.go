package distance

import (
	"encoding/json"
	"math"
	"testing"
)

func TestConversions(t *testing.T) {
	d := 1 * Kilometre

	testCases := []struct {
		name     string
		got      float64
		expected float64
	}{
		{"Metres", d.Metres(), 1000},
		{"Centimetres", d.Centimetres(), 100000},
		{"Millimetres", d.Millimetres(), 1000000},
		{"Kilometres", d.Kilometres(), 1},
		{"Miles", Mile.Miles(), 1},
		{"Yards", Mile.Yards(), 1760},
		{"Feet", Yard.Feet(), 3},
		{"Inches", Foot.Inches(), 12},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.expected {
				t.Errorf("Wrong conversion, got %v expected %v", tc.got, tc.expected)
			}
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		in       any
		expected string
	}{
		{"value", 5 * Metre, `"5m"`},
		{"pointer", func() *Distance { d := 5 * Metre; return &d }(), `"5m"`},
		{"struct field", struct{ D Distance }{2 * Kilometre}, `{"D":"2000m"}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := json.Marshal(tc.in)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if string(b) != tc.expected {
				t.Errorf("Wrong JSON, got %s expected %s", b, tc.expected)
			}
		})
	}
}

func TestUnmarshalJSON(t *testing.T) {
	testCases := []struct {
		in       string
		expected Distance
	}{
		{`"5m"`, 5 * Metre},
		{`"2km"`, 2 * Kilometre},
		{`"1760yd"`, Mile},
		{`null`, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			var d Distance
			if err := json.Unmarshal([]byte(tc.in), &d); err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if d != tc.expected {
				t.Errorf("Wrong distance, got %v expected %v", d, tc.expected)
			}
		})
	}
}

func TestUnmarshalJSONInvalid(t *testing.T) {
	var d Distance
	if err := json.Unmarshal([]byte(`"5furlong"`), &d); err == nil {
		t.Error("Expected error for unknown unit")
	}
}

func TestJSONRoundTrip(t *testing.T) {
	type payload struct {
		Run Distance `json:"run"`
	}

	in := payload{Run: 5 * Kilometre}
	b, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	var out payload
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if out != in {
		t.Errorf("Round trip mismatch, got %+v expected %+v", out, in)
	}
}

func TestJSONRoundTripExtremes(t *testing.T) {
	for _, d := range []Distance{1e-7, 1e21, -Mile, 0} {
		t.Run(d.String(), func(t *testing.T) {
			b, err := json.Marshal(d)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			var out Distance
			if err := json.Unmarshal(b, &out); err != nil {
				t.Fatalf("Unexpected error unmarshalling %s: %v", b, err)
			}
			if out != d {
				t.Errorf("Round trip mismatch, got %v expected %v", out, d)
			}
		})
	}
}

func TestMarshalJSONNonFinite(t *testing.T) {
	for _, d := range []Distance{Distance(math.NaN()), Distance(math.Inf(1)), Distance(math.Inf(-1))} {
		if _, err := json.Marshal(d); err == nil {
			t.Errorf("Expected error marshalling %v", d)
		}
	}
}

func TestUnmarshalJSONNumber(t *testing.T) {
	var d Distance
	if err := json.Unmarshal([]byte(`5`), &d); err == nil {
		t.Error("Expected error for JSON number")
	}
}
