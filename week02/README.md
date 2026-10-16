
# Base Converter CLI

A command-line tool written in Go that converts numbers between bases 2 and 16.

## Usage

```bash
go run . <number> <from-base> <to-base>
```

## Examples

Convert decimal to binary:

```bash
go run . 255 10 2
```

Output: `11111111`

Convert binary to decimal:

```bash
go run . 1010 2 10
```

Output: `10`

Convert octal to decimal:

```bash
go run . 77 8 10
```

Output: `63`

Convert hexadecimal to decimal:

```bash
go run . FF 16 10
```

Output: `255`

## Validation

The program rejects invalid digits, empty numbers, and bases outside 2 to 16.

## Tests

Run all tests:

```bash
go test ./...
```

## Implementation

The project uses custom `FromBase` and `ToBase` functions.
The `strconv` package is only used to parse command-line base arguments.
