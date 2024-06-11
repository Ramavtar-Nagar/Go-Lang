package main

import (
	"fmt"
)

func main(){

	name := "Oggy"
	age := 23
	height := 5.83293

	// Println = print line

	// it adds and executes the code in next line and also 
	// and also add the spaces in between of parameters

	fmt.Println("name : ", name, "age : ", age, "height : ", height); fmt.Println("Hello World")

	// output should be = name : oggyage : 23height : 5.83
	// but it comes as name :  Oggy age :  23 height :  5.83 coz space is added by Println function



	// Printf = print formatter
	// allows us to control the output format by using format specifiers
	fmt.Printf("age is %d\n", age)

	// fmt.Printf("Height is %d\n", height)
	// this does not work because %d format specifiers is for integers only but it is a float data type(that is Height)

	// so 
	fmt.Printf("Height is %.4f\n", height) // .4f gives only 4 digits after decimal
	fmt.Printf("Height is %.2f\n", height) // similarly it gives only 2 digits after decimal

	// it also moves the cursor to the next line
	// for example 
	fmt.Printf("Height is %.4f", height)
	fmt.Printf("Height is %.2f\n", height) // output: Height is 5.8329Height is 5.83

	// Printf is also used to find the data type of any particular variable
	fmt.Printf("Type of name is: %T\n", name)
	fmt.Printf("Type of age is: %T\n", age)
	fmt.Printf("Type of height is: %T\n", height)

	// multiple specifiers in a single statement
	fmt.Printf("Name: %s, Age: %d, Height: %.2f", name, age, height)

}