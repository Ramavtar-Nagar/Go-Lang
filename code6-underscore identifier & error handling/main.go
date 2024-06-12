package main

import (
	"fmt"
)

func divide(a, b float64) (float64, error){
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a/b, nil
}

func main() {

	fmt.Println("Error handling in Go-Lang")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	}else{
		fmt.Println("Divide Result:", result)
	}

	// if we want to ignore this err in that case we case use Underscore Identifier
	result, _ := divide(10, 2)
	fmt.Println("Divide Result:", result)
}