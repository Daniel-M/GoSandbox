package main

import(
	"fmt"
	"unsafe" // this makes this code unportable.
)

func main(){
	// Declaring an array of ints with 5 places
	// Each place is filled with 0.
	var x [8]int

	x[7] = 100

	fmt.Println("The array has", x)
	fmt.Printf("Size of %T is %v bytes\n", x, unsafe.Sizeof(x))
	
	fmt.Println("I can print array components, like this,",x[0])
	fmt.Println("I can print array components, like this,",x[7])

	// A shorter way to create arrays
	// Here i'm using auto casting
	y := [10]float32{
		1.0,
		2.0,
		3.0,
		4.0,
		5.0,
		6.0,
		7.0,
		8.0,
		9.0,
		10.0,
	}
	// this could be on a single line too but you already knew that, dont ya?


	// To create dynamically modifiable arrays we need to create slices
	var fSlice []float32
	
	//fSlice = make([]float32,y,10)
	// I'm telling to add from y[0] to y[9] to the slice
	// Be careful with the index numbering
	// tho is zero based, it needs "10" cuz is the total size
	fSlice = y[0:10]	

	fmt.Println("Original array",y)
	fmt.Println("A slice of the whole array",fSlice)

	// This should reduce the size of the array to a single place
	fSlice = make([]float32,1)
	
	fmt.Println("I have srinked the slice to",fSlice)
	
	fSlice = y[5:10]	
	fmt.Println("Now took a slice:",fSlice)
	
	fSlice = make([]float32,1)

	fmt.Println("Lets iterate through all the array making slices")

	for i:=1; i <= len(y) ; i++ {
	
		fSlice = y[0:i]	
		fmt.Println(i,"\t",fSlice)

	}
	for i:=len(y); i > 0 ; i-- {
	
		fSlice = y[0:i]	
		fmt.Println(i,"\t",fSlice)

	}
	fmt.Println("Fabulous, right?\nLets append y again to our slice");

	// Resize slice to store two times the capacity of array y
	fSlice = make([]float32,2*len(y))

	// This copies y into fist 10 places of fSlice
	copy(fSlice,y[0:10])
	// This copies y into the remaining 10 places of fSlice
	// That's why i've choosen fSlice[10:20] the last 10
	copy(fSlice[10:20],y[0:10])

	fmt.Println(fSlice)

}
