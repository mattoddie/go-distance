# go-distance

[![Test](https://github.com/mattoddie/go-distance/actions/workflows/test.yaml/badge.svg)](https://github.com/mattoddie/go-distance/actions/workflows/test.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/mattoddie/go-distance/v2.svg)](https://pkg.go.dev/github.com/mattoddie/go-distance/v2)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/mattoddie/go-distance/badge)](https://scorecard.dev/viewer/?uri=github.com/mattoddie/go-distance)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A Go package for representing distances, converting between metric and
imperial units, and parsing distance strings in the style of
`time.ParseDuration`.

## Installation

```sh
go get github.com/mattoddie/go-distance/v2
```

Requires Go 1.22 or later. The package has no dependencies outside the
standard library.

## Usage

```go
import "github.com/mattoddie/go-distance/v2"

d, err := distance.Parse("1mi200yd")
if err != nil {
	// handle error
}
fmt.Printf("%.2f km\n", d.Kilometres()) // 1.79 km

marathon := 42.195 * distance.Kilometre
fmt.Printf("%.1f miles\n", marathon.Miles()) // 26.2 miles
fmt.Println(marathon)                        // 42195m
```

### Units

| Unit       | Suffix | Constant     | Method          |
| ---------- | ------ | ------------ | --------------- |
| Millimetre | `mm`   | `Millimetre` | `Millimetres()` |
| Centimetre | `cm`   | `Centimetre` | `Centimetres()` |
| Metre      | `m`    | `Metre`      | `Metres()`      |
| Kilometre  | `km`   | `Kilometre`  | `Kilometres()`  |
| Inch       | `in`   | `Inch`       | `Inches()`      |
| Foot       | `ft`   | `Foot`       | `Feet()`        |
| Yard       | `yd`   | `Yard`       | `Yards()`       |
| Mile       | `mi`   | `Mile`       | `Miles()`       |

`Parse` accepts a possibly signed sequence of numbers, each with an optional
fraction and a unit suffix, such as `"300ft"`, `"-1.5km"` or `"1mi200yd"`.

### JSON

`Distance` implements `json.Marshaler` and `json.Unmarshaler`. It encodes as a
string in metres and decodes from any string `Parse` accepts:

```go
type Run struct {
	Length distance.Distance `json:"length"`
}

b, _ := json.Marshal(Run{Length: 5 * distance.Kilometre})
fmt.Println(string(b)) // {"length":"5000m"}

var r Run
_ = json.Unmarshal([]byte(`{"length":"3.1mi"}`), &r)
```

## Upgrading from v1

v2 contains breaking changes:

- Import `github.com/mattoddie/go-distance/v2` instead of
  `github.com/mattoddie/go-distance`.
- `ParseDistance` is now `Parse`.
- `Distance` values always marshal to JSON as strings such as `"5m"`. In v1 a
  `Distance` value marshalled as a bare number, which could not be
  unmarshalled again. If you stored v1 JSON containing numbers, convert those
  values to strings before decoding them with v2.
- `String` never uses exponent notation (`"0.0000001m"` rather than
  `"1e-07m"`), so every finite distance round trips through `Parse`.

See the [changelog](CHANGELOG.md) for the full list.

## Contributing

Contributions are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md) to get
started, and please follow the [Code of Conduct](CODE_OF_CONDUCT.md). To report
a security issue, see [SECURITY.md](SECURITY.md).

## License

go-distance is released under the [MIT License](LICENSE).

The parser is adapted from the duration parser in the Go standard library's
`time` package, which is distributed under a BSD-style licence reproduced in
[LICENSE-GO](LICENSE-GO).
