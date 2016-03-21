package main

import (
	"fmt"
	"unsafe" // This is for the SizeOf function. This code is not portable tho.
)

func main(){

	/* Basic variable types in Go */

	// Boolean type
	var bool_t bool

	/* Unsigned int types */

	// Unsigned int of 8 bits
	var uint8_t uint8
	// Unsigned int of 16 bits
	var uint16_t uint16
	// Unsigned int of 32 bits
	var uint32_t uint32
	// Unsigned int of 64 bits
	var uint64_t uint64
	// standard unsigned int, its size is architecture dependent (32 or 64 bits)
	var uint_t uint

	/* Singed int */

	// 8 bits int
	var int8_t int8
	// 16 bits int
	var int16_t int16
	// 32 bits int
	var int32_t int32
	// 64 bits int
	var int64_t int64
	// standard int, its size is architecture dependent (32 or 64 bits)
	var int_t int

	/* Special types architecture independent */

	// 8 bits int, aka a byte.
	// shorthand for uint8
	var byte_t byte
	// 32 bits int that can be used to iterate through UTF-8 characters
	var rune_t rune

	/* Floating point types */

	// 32 bit float
	var float32_t float32
	// 64 bit float
	var float64_t float64

	/* Complex numbers native support */

	// 32 bit real part, 32 bit imaginary part
	var complex64_t complex64
	// 64 bit real part, 64 bit imaginary part
	var complex128_t complex128


	/* String types */

	var y string
	// String type initialized to a word
	var x string = "Hello World"


	/* Multiple declaration of variales */

	// Declare several variables at once using blocks 
	// commented bc wont compile cuz variables are unused
/*	var (
		int8_t_1 int8
		int8_t_2 int16
		int8_t_3 int32
		int8_t_4 int64
		float32_t_1 float32
		complex64_t_1 complex64
	)
*/

	/* Declaring constant values throughout the code */

	// Declare constant values
	const c_int8_t=1

	// Declare several constant values at once using a block
	const (
		c_int8_t_1 = 1
		c_int8_t_2 = 2
		c_int8_t_3 = 3
		c_int8_t_4 = 4
		c_float32_t_1 = 5
		c_complex64_t_1 = 1
	)


	/* Multiple value assignation at declaration time */

	// a = 1, b = 2, c = 3
	// commented bc wont compile cuz variables are unused
/*	var a,b,c int = 1, 2, 3
	int8_t_1 = 0
	int8_t_2 = 2
	int8_t_3 = 3
	int8_t_4 = 4
	float32_t_1 = 0 
	complex64_t_1 = 0
*/

	y = "Hello World"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x+y)
	y=x+y
	fmt.Println(y)

	fmt.Printf("Size of %T is %v bytes\n",bool_t,unsafe.Sizeof(bool_t))
	fmt.Printf("Size of %T is %v bytes\n",uint8_t,unsafe.Sizeof(uint8_t))
	fmt.Printf("Size of %T is %v bytes\n",uint16_t,unsafe.Sizeof(uint16_t))
	fmt.Printf("Size of %T is %v bytes\n",uint32_t,unsafe.Sizeof(uint32_t))
	fmt.Printf("Size of %T is %v bytes\n",uint64_t,unsafe.Sizeof(uint64_t))
	fmt.Printf("Size of %T is %v bytes\n",uint_t,unsafe.Sizeof(uint_t))
	fmt.Printf("Size of %T is %v bytes\n",int8_t,unsafe.Sizeof(int8_t))
	fmt.Printf("Size of %T is %v bytes\n",int16_t,unsafe.Sizeof(int16_t))
	fmt.Printf("Size of %T is %v bytes\n",int32_t,unsafe.Sizeof(int32_t))
	fmt.Printf("Size of %T is %v bytes\n",int64_t,unsafe.Sizeof(int64_t))
	fmt.Printf("Size of %T is %v bytes\n",int_t,unsafe.Sizeof(int_t))
	fmt.Printf("Size of %T is %v bytes\n",byte_t,unsafe.Sizeof(byte_t))
	fmt.Printf("Size of %T is %v bytes\n",rune_t,unsafe.Sizeof(rune_t))
	fmt.Printf("Size of %T is %v bytes\n",float32_t,unsafe.Sizeof(float32_t))
	fmt.Printf("Size of %T is %v bytes\n",float64_t,unsafe.Sizeof(float64_t))
	fmt.Printf("Size of %T is %v bytes\n",complex64_t,unsafe.Sizeof(complex64_t))
	fmt.Printf("Size of %T is %v bytes\n",complex128_t,unsafe.Sizeof(complex128_t))
	fmt.Printf("Size of %T is %v bytes\n",y,unsafe.Sizeof(y))
	
	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("1.0 + 1.0 =", 1.0 + 1.0)
	fmt.Println("len(\"Hello World\": ")
	fmt.Println(len("Hello World"))
	fmt.Println("\"Hello World[1]\"")
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	fmt.Println(" Boolean logic:")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
