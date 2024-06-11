package main

import (
	"fmt"
)

	// Basic Function Syntax
	// --->
	/** In Go, functions are defined using the func keyword followed by the 
	function name, parameters, return types, and the function body. **/
	
	// Function definition
	func add(a int, b int) int {
		return a+b
	}



	// Parameters and Return Types
	// --->
	/** Parameters: Functions can take zero or more parameters. 
	Parameters are defined within the parentheses after the function name.**/
	// --->
	/** Return Types: Functions can return zero or more values. 
	Return types are defined after the parameters. **/

	// Examples =>
	func multiply(a int, b int) int {
		return a * b
	}

	func divide(a, b float64) (float64, error){
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}

		return a/b, nil
	}


	// Named Return Values
	// ---> 
	/** Go allows you to name the return values in the function signature. 
	This can make your code more readable and allows you to use a `return` statement
	without specifying return values. **/

	// example =>
	func swap(a, b string) (x, y string){
		a = y
		b = x
		return
	}

	// Multiple Return Values
	/** A function in Go can return multiple values. This feature is often used for error handling.
	Like we did in the above divide function declaration **/

	// func divide(a, b float64) (float64, error) {
	// 	if b == 0 {
	// 		return 0, fmt.Errorf("cannot divide by zero")
	// 	}

	// 	return a/b, nil
	// }

	// Variadic Functions
    /** Variadic functions can accept a variable number of arguments. The syntax for a 
	variadic parameter uses an ellipsis (...) followed by the type. **/

	func sum(nums ...int) int {
		total := 0
		for _, num := range nums {
			total += num 
		}
		return total
	}

	// Anonymous Functions and Closures
	/** Go supports anonymous functions, which can be used as closures. Anonymous functions
	 are useful when you need a quick function without declaring it beforehand. They are declared 
	 and used in main function only **/



	 // Defer Statement
    /** The defer statement is used to ensure that a function call is performed later in a program’s
	 execution, usually for purposes of cleanup. Deferred function calls are executed in LIFO (last-in-first-out) 
	 order. **/

	 func main(){
		defer fmt.Println("World")
		fmt.Println("Hello")
	 }


	 // Function Types and Higher-Order Functions
	 /** Functions are first-class citizens in Go, meaning you can assign them to variables, pass them as arguments, 
	 and return them from other functions. **/


	 /** we can not call this function in file because in a file only one main() function is allowed to be executed **/
	 func mainb(){
		// Assigning a function to a variable
		/**
			f := func(a, b int) int {
				return a + b
			}
			fmt.Println(f(3,4))

			// Passing a function as an argument
			passingFunction := func(f func(int, int) int, a int, b int ) int {
				return(a, b)
			}
			fmt.Println(passingFunction(f, 10, 20))

			// Returning a function from another function
			multiplier := func(x int) func(int) int {
				return func(y int) int {
					return x * y
				}
			}
			double := multiplier(3)
			resultOfTwoFunctions := double(11)
			fmt.Println("Result of function inside another function as result is: ", resultOfTwoFunctions)
		**/

	 }

func main() {

	fmt.Println("We are learning Functions in Go-Lang")
	sum := add(3,6)
	fmt.Println("The sum of your values is: ", sum)

	multiplyResult := multiply(3, 3)
	fmt.Println("Multiply Result:", multiplyResult)

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	}else{
		fmt.Println("Divide Result:", result)
	}

	sumResult := sum(1,2,3,4,5,6)
	fmt.Println("Total sum is", sumResult)

	// Anonymous Functions 
	additionresult := func addition(a, b int) int {
		return a + b
	}
	fmt.Println(additionresult(3,6))

	 // Closure
	 x := 10
	 increament := func() int {
		x++
		return x
	 }
	 fmt.Println(increament()) // 11
	 fmt.Println(increament()) // 12

}