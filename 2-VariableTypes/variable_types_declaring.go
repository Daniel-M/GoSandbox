package main

import "fmt"

func main(){
	var x string = "Hello World"
	var y string
	y = "Hello World"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x+y)
	y=x+y
	fmt.Println(y)
}
