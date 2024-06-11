package main

import (
	"fmt"
)

func main(){
	fmt.Println("Learning Go Language --->");

	var name string = "Random Access Memory"
	fmt.Println(name)
	
	var version = 36
	fmt.Println(version)

	var money int = 30000
	var currency = 83.6
	fmt.Println(money)
	fmt.Println("Currency in dollars rn : ", currency)

	var diameter float64 = 16.9
	fmt.Println(diameter)

	var decision bool = true
	fmt.Println(decision)

	var student = "Aman"
	fmt.Println(student)

	// variable declarations using [const] which means that the data can never be changed
	const pi = 3.14
	// pi = 30
	fmt.Println(pi)

	// whereas in case of declarations using var
	name = "RAM"
	fmt.Println(name)

	// variables declarations shorthand without specifying var OR const
	student2 := "joshua"
	fmt.Println(student2)

	// declarations of variables with visibility outside the package by capitalization of its name
	// public variable (exported)
	var publicVariable = "This is a public variable"
	fmt.Println(publicVariable)

	// private variable (unexported)
	var privateVariable = "This is a private private"
	fmt.Println(privateVariable)

	similar approach for functions
	func PublicFunction() {
		fmt.Println("This is a public function")
	}

	func privateFunction() {
		fmt.Println("This is a private function")
	}

	

}
