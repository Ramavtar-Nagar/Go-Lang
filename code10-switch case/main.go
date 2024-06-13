package main

import(
	"fmt"
)

func main(){

	fmt.Println("Switch Case in go -lang --->")


	// Normal switch case statement
	day := "monday"

	switch day {
	case "monday":
		fmt.Println("Today is monday")
	case "tuesday":
		fmt.Println("Today is Tuesday")
	case "wednesday":
		fmt.Println("Today is Wednesday")
	case "thursday":
		fmt.Println("Today is Thursday")
	case "friday":
		fmt.Println("Today is Friday")
	case "saturday":
		fmt.Println("Today is Saturday")
	case "sunday":
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Unknown day - fun day")
	}



	// Switch case statement with multiple values 
	month := "february"
	switch month{
	case "february", "march":
		fmt.Println("This is season on FALL")
	case "aprail", "may", "june":
		fmt.Println("This is season on Summer")
	case "july", "august", "september", "october":
		fmt.Println("This is season on Rain")
	case "november", "december", "january":
		fmt.Println("This is season on Winter")
	default:
		fmt.Println("Other season")
	}


	// switch case statement with expressions

	temperature := 13

	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cool")
	case temperature >= 20 && temperature < 35:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}

}