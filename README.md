# goconv

A lightweight Go library for effortless type conversion — no errors, just sensible defaults.

## Description

`goconv` simplifies type conversions in Go by providing a set of intuitive functions that never return errors. Instead of dealing with error handling for every conversion, goconv returns sensible default values (like 0, empty string, or false) when conversions fail. This makes your code cleaner and more readable, especially when working with dynamic data sources like JSON, user input, or configuration files.

**Key Features:**
- Zero error handling required
- Sensible defaults for failed conversions
- Support for all common Go types (int, uint, float, string, bool)
- Safe negative-to-unsigned conversions (clamps to 0)
- Consistent API across all conversion functions

## Installation

```bash
go get github.com/dharper7051/goconv
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/dharper7051/goconv"
)

func main() {
    // String conversions
    fmt.Println(goconv.String(42))           // "42"
    fmt.Println(goconv.String(3.14159))      // "3.141590" (6 decimal places)
    fmt.Println(goconv.String(true))         // "true"
    fmt.Println(goconv.String(nil))          // ""

    // Integer conversions
    fmt.Println(goconv.Int("123"))           // 123
    fmt.Println(goconv.Int("invalid"))       // 0
    fmt.Println(goconv.Int(3.14))            // 3
    fmt.Println(goconv.Int64(42))            // 42

    // Unsigned integer conversions (negatives clamp to 0)
    fmt.Println(goconv.Uint(-5))             // 0
    fmt.Println(goconv.Uint("100"))          // 100
    fmt.Println(goconv.Uint64("999"))        // 999

    // Float conversions
    fmt.Println(goconv.Float64("3.14"))      // 3.14
    fmt.Println(goconv.Float32(42))          // 42.0
    fmt.Println(goconv.Float64("invalid"))   // 0.0

    // Boolean conversions
    fmt.Println(goconv.Bool(1))              // true
    fmt.Println(goconv.Bool(0))              // false
    fmt.Println(goconv.Bool("true"))         // true
    fmt.Println(goconv.Bool("false"))        // false
    fmt.Println(goconv.Bool("invalid"))      // false
}
```

### Available Functions

- `String(value interface{}) string` - Convert any value to string
- `Int(value interface{}) int` - Convert to int
- `Int32(value interface{}) int32` - Convert to int32
- `Int64(value interface{}) int64` - Convert to int64
- `Uint(value interface{}) uint` - Convert to uint (negatives become 0)
- `Uint32(value interface{}) uint32` - Convert to uint32 (negatives become 0)
- `Uint64(value interface{}) uint64` - Convert to uint64 (negatives become 0)
- `Float32(value interface{}) float32` - Convert to float32
- `Float64(value interface{}) float64` - Convert to float64
- `Bool(value interface{}) bool` - Convert to bool

### Default Behavior

When conversions fail or receive invalid input, goconv returns sensible defaults:
- Numeric types: `0`
- String: `""` (empty string)
- Bool: `false`

This eliminates the need for error checking while maintaining predictable behavior.
