package main

import (
	"fmt"

	"github.com/Eishanfaust/go-tests/lib1"
)

func main() {
	sum := lib1.Add(3, 5)
	product := lib1.Multiply(4, 6)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)
}
