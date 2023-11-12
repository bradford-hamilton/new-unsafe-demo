package main

import (
	"fmt"
	"unsafe"

	"github.com/bradford-hamilton/new-unsafe-demo/internal/priv"
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

	fruits := []string{"apples", "oranges", "bananas", "kansas"}
	start := unsafe.Pointer(unsafe.SliceData(fruits))
	size := unsafe.Sizeof(fruits[0])

	for i := 0; i < len(fruits); i++ {
		p(*(*string)(unsafe.Add(start, uintptr(i)*size)))
	}
	// apples
	// oranges
	// bananas
	// kansas

	// ------------------------------------------------
	ps := priv.NewStruct()
	p(ps) // {foo:bar bar:1337 baz:[100 150 200 250]}

	foo := (*string)(unsafe.Pointer(&ps))
	*foo = "bradford"
	p(ps) // {foo:bradford bar:1337 baz:[100 150 200 250]}

	bar := (*int)(unsafe.Add(unsafe.Pointer(&ps), unsafe.Sizeof("")))
	*bar = 20
	p(ps) // {foo:bradford bar:20 baz:[100 150 200 250]}

	slcPtr := (*[]int)(unsafe.Add(
		unsafe.Pointer(&ps), (unsafe.Sizeof("") + unsafe.Sizeof(int(0))),
	))
	p(*slcPtr) // [100 150 200 250]

	start = unsafe.Pointer(unsafe.SliceData(*slcPtr))
	size = unsafe.Sizeof(int(0))

	for i := 0; i < len(*slcPtr); i++ {
		p(*(*int)(unsafe.Add(start, uintptr(i)*size)))
	}
	// 100
	// 150
	// 200
	// 250
}

func p(a any) { fmt.Printf("%+v\n", a) }
