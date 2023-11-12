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
}

func p(a any) { fmt.Printf("%+v\n", a) }
