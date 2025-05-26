package calc_test

import (
	"fmt"

	"github.com/ezhdanovskiy/calc-lib-example"
)

func ExampleAdd() {
	result := calc.Add(5, 3)
	fmt.Println(result)
	// Output: 8
}

func ExampleAdd_negative() {
	result := calc.Add(-10, 5)
	fmt.Println(result)
	// Output: -5
}

func ExampleAdd_zero() {
	result := calc.Add(0, 0)
	fmt.Println(result)
	// Output: 0
}