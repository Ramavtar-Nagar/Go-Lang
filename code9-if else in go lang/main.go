package main

import(
	"fmt"
)

func main(){

	fmt.Println("If - Else in go lang --->")

	x := 23
	if x > 5{
		fmt.Println("x is greater than 5")
	} else{
		fmt.Println("x is smaller than 5")
	}

	z := 6
	if z > 10{
		fmt.Println("z is greater than 10")
	}else if z > 5{
		fmt.Println("z is greater than 5 but smaller than 10")
	}else{
		fmt.Println("z is smaller than 5")
	}

	// with condition of && and || 
	y := 10
	if x > 5 && y > 5{
		fmt.Println("x and y both are greater than 5")
	}else{
		fmt.Println("one from x or y is smaller than 5")
	}

	// with condition of && and || over another condition
	if x > 5 && (y > 5 || z < 30){
		fmt.Println("Hey how are you ?")
	}else{
		fmt.Println("condition over conditions hehe...")
	}
}