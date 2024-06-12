package main

import (
	"fmt"
	"slices"
)

func main(){
	
	fmt.Println("Slices in Go-Lang --->")

	// initializing slice
	number := []int{1,2,3,4,5,6}
	fmt.Println("Number slice values are:", number)
	fmt.Printf("Number has data type: %T\n", number)
	fmt.Println("Length of nnumbers slice is:", len(number))

	// to append something in slice
	number = append(number, 6, 7, 8, 9, 10, 11, 12, 13)
	fmt.Println("Number slice values are:", number)
	fmt.Println("Length of nnumbers slice is:", len(number))

	// initializing empty slice
	name := []string{}
	fmt.Println("name:", name)

	// slice in deep
	rollNum := []int{1,2,3}
	fmt.Println("Values of Slice:", rollNum)
	fmt.Println("Length of slice", len(rollNum))
	fmt.Println("Capacity of slice:",cap(rollNum))


	/** initializing slice with make method to set Length and 
	Capacity of slice according to us **/

	students := make([]string, 3,5)
	fmt.Println("Values of students Slice:", students) // ["" "" ""]
	fmt.Println("Length of students slice", len(students)) // 3
	fmt.Println("Capacity of students slice:",cap(students)) // 5

	students = append(students, "Siya")
	fmt.Println("Values of students Slice:", students) // ["" "" "" Siya]
	fmt.Println("Length of students slice", len(students)) // 4
	fmt.Println("Capacity of students slice:",cap(students)) // 5

	students = append(students, "Ram")
	fmt.Println("Values of students Slice:", students) // ["" "" "" Siya Ram]
	fmt.Println("Length of students slice", len(students)) // 5
	fmt.Println("Capacity of students slice:",cap(students)) // 5
	
	/** if we append one more element after the length of the slice than length 
	increases by the no. of elements appended and capacity becomes doube **/
	students = append(students, "Jay")
	fmt.Println("Values of students Slice:", students) // ["" "" "" Siya Ram jai]
	fmt.Println("Length of students slice", len(students)) // 6
	fmt.Println("Capacity of students slice:",cap(students)) // 10

}