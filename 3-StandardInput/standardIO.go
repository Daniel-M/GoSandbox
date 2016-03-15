package main

import "fmt"

func main(){
	var fNumber float64

	fmt.Println("Standard Input from keyboard")
	fmt.Println("Enter a number:")
	fmt.Scanf("%f",&fNumber)

	fmt.Println("You entered the number ",fNumber)
	fmt.Println("The number you entered times 2 = ",2*fNumber)
	
}
