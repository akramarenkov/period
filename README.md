# Period

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/period.svg)](https://pkg.go.dev/github.com/akramarenkov/period)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/period)](https://goreportcard.com/report/github.com/akramarenkov/period)
[![codecov](https://codecov.io/gh/akramarenkov/period/releases/tag/v0.1.0/badge.svg?token=2E4F42B30C)](https://codecov.io/gh/akramarenkov/period)

## Purpose

Library that extends time.Duration from standard library with years, months and days

Without approximations

Compatible with time.Duration

## Usage

Example:

```go
package main

import (
    "fmt"
    "time"

    "github.com/akramarenkov/period"
)

func main() {
    period, found, err := period.Parse("2y3mo10d23h59m58s10ms30µs10ns")
    if err != nil {
        panic(err)
    }

    if !found {
        return
    }

    fmt.Println(period, period.ShiftTime(time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)))
    // Output: 2y3mo10d23h59m58.01003001s 2025-07-11 23:59:58.01003001 +0000 UTC
}
```
