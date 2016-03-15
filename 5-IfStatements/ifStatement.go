package main

import "fmt"

func main(){
	var fNumber int 
	var i int

	fmt.Println("Standard Input from keyboard")
	fmt.Print("Enter a number for the loop: ")
	fmt.Scanf("%d",&fNumber)

	// loop variable created with automatic type detection
	for i = 0; i < fNumber; i++ {

		// The else must be written like that or will send compilation errors
		// "like that" i.e. that ugly way of } else
		if i % 2 == 0 {
			fmt.Println(i,"even")
		} else {
			fmt.Println(i,"odd")
		}

	}

	// How to make successive if, else if, else statements
	if i % 2 == 0 {
	// divisible by 2
	} else if i % 3 == 0 {
	// divisible by 3
	} else if i % 4 == 0 {
	// divisible by 4
	}
	
	switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		default: fmt.Println("Unknown Number")
	}

}
