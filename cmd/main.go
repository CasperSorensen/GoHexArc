package main

import (
	"GoHexArc/internal/adapters/core/arithmetic"
	"GoHexArc/internal/ports"
	"fmt"
)

func main() {
	fmt.Println("Module running")

	// ports
	var core ports.ArithmeticPort

	core = arithmetic.NewAdapter()

	arithAdapter := arithmetic.NewAdapter()
	result, err := arithAdapter.Addition(2, 3)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}
