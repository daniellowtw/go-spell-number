# go-spell-number

A small dependency free library to turn spell out numbers. It supports positive numbers up to uint64.

# Usage
```go
package main

import (
    "github.com/daniellowtw/go-spell-number"
    "math"
)

func main() {
    spell.Number(123) // one hundred twenty three
    spell.Number(4123) // four thousand one hundred twenty three
    spell.Number(1e6) // one million
    spell.Number(math.MaxInt64) // nine quintillion ...
    spell.LargeNumber(math.MaxUint64) // eighteen quintillion ...
}
```