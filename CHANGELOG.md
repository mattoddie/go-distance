# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Breaking changes

- The module path is now `github.com/mattoddie/go-distance/v2`.
- `ParseDistance` has been renamed to `Parse`.
- `Distance` now always marshals to JSON as a string such as `"5m"`.
  Previously a `Distance` value marshalled as a bare number, which could not
  be unmarshalled again, and a `*Distance` produced invalid JSON.
- `String` no longer uses exponent notation, so very small or large distances
  are written out in full (for example `"0.0000001m"` instead of `"1e-07m"`).
- `MarshalJSON` returns an error for NaN and infinite distances.
- Parse error messages now quote the input, for example
  `distance: unknown unit "furlong" in distance "5furlong"`.

### Added

- Numbers of any length are accepted by `Parse`, rather than at most 18
  digits on each side of the decimal point.
- `Parse` returns an error instead of an infinite distance when the result
  overflows.
- Package documentation, examples and a fuzz test for the parser.
- MIT licence, contributing guide, code of conduct and security policy.

### Fixed

- JSON round trips for all finite distances, including very small and very
  large values.

## [1.1.0] - 2019-10-03

- Add Go modules support.

## [1.0.1] - 2019-03-13

- Add support for marshalling and unmarshalling JSON.

## [1.0.0] - 2014-06-18

- Initial release.

[Unreleased]: https://github.com/mattoddie/go-distance/compare/v1.1.0...HEAD
[1.1.0]: https://github.com/mattoddie/go-distance/compare/v1.0.1...v1.1.0
[1.0.1]: https://github.com/mattoddie/go-distance/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/mattoddie/go-distance/releases/tag/v1.0.0
