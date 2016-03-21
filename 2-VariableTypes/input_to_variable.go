package main

import "fmt"

func main(){
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	// Automatic type assignation through operator :=
	output := input * 2

	fmt.Println(output)
}
