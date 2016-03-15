package main

import "fmt"

func main(){
	var fNumber int 

	fmt.Println("Standard Input from keyboard")
	fmt.Print("Enter a number for the loop: ")
	fmt.Scanf("%d",&fNumber)

	// loop variable created with automatic type detection
	for i := 0; i < fNumber; i++ {
		fmt.Println(i)
	}

	fmt.Println("Other way to make loops")

	var i int

	// loop variable created prior for loop
	for i=0; i<fNumber;i++{
		fmt.Println(i)
	}

	fmt.Println("Another way to make loops")

	i=0

	// for executes as long as the boolean expression evaluates to true.
	// Iteration variable is iterated inside the loop cycle.
	for  i<fNumber {
		fmt.Println(i)
		i++
	}

}
