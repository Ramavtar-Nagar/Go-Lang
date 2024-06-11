package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){


	// it scans till space using Scan method
	var name string
	var age int

	fmt.Println("Enter your name: ")
	fmt.Scan(&name) 
	/** if we give something after space character also (like more than one word) 
	than in that case it stops scanning here and does not scan anything after this
	 line of code **/
	
	fmt.Println("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println("Hello mr.", name, "Your age is:", age)
	// OR
	fmt.Printf("Hello Mr. %s, You are %d years old.", name, age)



	/** same is the case in Scanln method too, that if we give something after 
	space character also (like more than one word) than in that case it stops
	scanning here and does not scan anything after this line of code **/
	var schoolName string
	var class int

	fmt.Println("Enter your school name: ")
	fmt.Scanln(&schoolName)
	
	fmt.Println("Enter your class: ")
	fmt.Scanln(&class)

	fmt.Println("Your school is", schoolName, "Your class is:", class)
	// OR
	fmt.Printf("Your school is %s, You are in %d standard.", schoolName, class)



	/** to over this problem and scan the whole line for tasks such as reading 
	an entire line of text we have to do scanning with the help of bufio package **/
	
	// definying a new reader
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Tell me your father's name: ")
	fathersName, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(fathersName)  // Remove the newline character

	fmt.Print("What is your father's age: ");
	fathersAge, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(fathersAge)  // Remove the newline character

	fmt.Println("Your fathers name is:", fathersName, "and your fathers's age is:", fathersAge)
	
	// fmt.Printf("Your fathers name is %s, Your father is %d years old.", fathersName, fathersAge) 
	/** this approach does not work in this case, so to do that we have to 
	do this => name = strings.TrimSpace(fathersAge) **/
	
}