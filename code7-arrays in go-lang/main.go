package main

import (
	"fmt"
)

func main(){

	fmt.Println("Arrays in go lang --->")

	// as indexing starts from zero so here we get 0, 1, 2, 3, 4
	var studentsArray[5] string
	studentsArray[0] = "Aakash"
	studentsArray[2] = "siya"
	fmt.Println("Names of students is:", studentsArray)

	// arrays with initialization
	var numbers = [9]int{1,2,3,4,5,6}
	fmt.Println("Number are :", numbers)

	// Length of"" arrays
	fmt.Println("Length of studentsArray is:", len(studentsArray))
	fmt.Println("Length of numbers array is:", len(numbers))

	// Accessing values of array elements
	fmt.Println("Value of studentsArray array at 2nd index is:", studentsArray[2])
	fmt.Println("Value of numbers array at 2nd index is:", numbers[2])

	// Array initialization default values
	var price[5]int
	var values[5]float32
	var goods[5]string
	var flags[5]bool
	fmt.Println("Default values of price array:", price)
	fmt.Println("Default values of values array is:", values)
	fmt.Println("Default values of goods array is:", goods)
	fmt.Printf("Default values of goods array with formatter is: %q\n", goods)
	fmt.Println("Default values of flags array is:",flags)
}