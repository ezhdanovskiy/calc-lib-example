# calc-lib-example

A simple Go library for arithmetic operations.

## Installation

```bash
go get github.com/ezhdanovskiy/calc-lib-example
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/ezhdanovskiy/calc-lib-example"
)

func main() {
    result := calc.Add(5, 3)
    fmt.Println(result) // Output: 8
}
```

## Available Functions

### Add
`Add(a, b int) int` - Returns the sum of two integers.

## Examples

See the [example_test.go](example_test.go) file for more usage examples.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.