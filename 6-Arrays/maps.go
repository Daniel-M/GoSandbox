package main

import(
	"fmt"
	"unsafe" // this makes this code unportable.
)

func main(){
	// Maps work like c++ maps, you use a key to map stuff 
	// Here maps are declared
	var mapint map[int]int
	var mapstring map[string]int

	fmt.Printf("Size of empty %T is %v bytes\n",mapint,unsafe.Sizeof(mapint))
	fmt.Printf("Size of empty %T is %v bytes\n",mapstring,unsafe.Sizeof(mapstring))

	//Here maps are initialized :|
	mapint = make(map[int]int)
	mapstring = make(map[string]int)

	mapint[1] = 5
	mapint[2] = 4
	mapint[3] = 3
	mapint[4] = 2
	mapint[5] = 1

	mapstring["hey"] = 1
	mapstring["abc"] = 2
	mapstring["def"] = 3
	mapstring["hij"] = 4
	mapstring["klm"] = 5

	fmt.Println(mapint)
	fmt.Println(mapstring)
	fmt.Println("You can access single elements like",mapstring["klm"])

	fmt.Printf("Size of empty %T is %v bytes\n",mapint,unsafe.Sizeof(mapint))
	fmt.Printf("Size of empty %T is %v bytes\n",mapstring,unsafe.Sizeof(mapstring))
}
