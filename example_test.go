package distance_test

import (
	"fmt"

	"github.com/mattoddie/go-distance/v2"
)

func ExampleParse() {
	d, err := distance.Parse("1mi200yd")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.2f km\n", d.Kilometres())
	// Output: 1.79 km
}

func ExampleDistance_Miles() {
	marathon := 42.195 * distance.Kilometre
	fmt.Printf("%.1f miles\n", marathon.Miles())
	// Output: 26.2 miles
}
