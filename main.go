package main

import (
	"fmt"
	"unsafe"
)

type user struct {
	name    string
	age     int
	animals []string
}

func main() {
	var u user
	p(u) // {name: age:0 animals:[]}

	uNamePtr := (*string)(unsafe.Pointer(&u))
	*uNamePtr = "bradford"
	p(u) // {name:bradford age:0 animals:[]}

	age := (*int)(unsafe.Add(unsafe.Pointer(&u), unsafe.Offsetof(u.age)))
	*age = 34
	p(u) // {name:bradford age:34 animals:[]}

	u.animals = []string{"missy", "ellie", "toby"}
	secondAnimal := (*string)(unsafe.Add(
		unsafe.Pointer(unsafe.SliceData(u.animals)),
		unsafe.Sizeof(""),
	))
	p(u) // {name:bradford age:34 animals:[missy ellie toby]}

	*secondAnimal = "carlos"
	p(u) // {name:bradford age:34 animals:[missy carlos toby]}

	// ------------------------------------------------

	mySlice := []string{"apples", "oranges", "bananas", "kansas"}
	start := unsafe.Pointer(unsafe.SliceData(mySlice))
	size := unsafe.Sizeof(mySlice[0])
	for i := 0; i < len(mySlice); i++ {
		p(*(*string)(unsafe.Add(start, uintptr(i)*size)))
	}
}

func p(a any) { fmt.Printf("%+v\n", a) }
