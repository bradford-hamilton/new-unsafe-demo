package main

import (
	"fmt"
	"unsafe"

	"github.com/bradford-hamilton/new-unsafe-demo/internal/priv"
)

func p(a any) { fmt.Printf("%+v\n", a) }

type user struct {
	name    string
	age     int
	animals []string
}

func main() {
	// Declare zero value 'user' struct and print its contents:
	var u user
	p(u) // {name: age:0 animals:[]}

	// Retrieve an unsafe.Pointer to 'u', which points to the first
	// member of the struct - 'name' - which is a string. Then we
	// cast the unsafe.Pointer to a string pointer. This allows us
	// to manipulate the memory pointed at as a string type.
	uNamePtr := (*string)(unsafe.Pointer(&u))
	*uNamePtr = "bradford"
	p(u) // {name:bradford age:0 animals:[]}

	// Here we have a similar situation in that we want to get a pointer
	// to a struct member. This time it is the second member, so we need
	// to calculate the address within the struct by using offsets. The
	// general idea is that we need to add the size of 'name' to the
	// address of the struct to get to the start of the 'age' member.
	// Finally we cast it to an '*int' - so we get an unsafe.Pointer
	// from 'unsafe.Add' and cast it to an '*int'.
	age := (*int)(unsafe.Add(unsafe.Pointer(&u), unsafe.Offsetof(u.age)))
	*age = 34
	p(u) // {name:bradford age:34 animals:[]}

	// Note: this can be applied to private structs/struct members,
	// however there are some implementation differences. For example,
	// when getting a pointer to the 'age' member, we wouldn't have
	// access to 'u.age'. Instead, we could use unsafe.Sizeof("") to
	// get the size of a string (the first member) and add that to the
	// base 'u' unsafe.Pointer.

	// Here we are working with something a bit different. First we add
	// a slice of animals to the user struct we've been working with.
	u.animals = []string{"missy", "ellie", "toby"}

	// Now we want to get a pointer to the second slice element and make
	// a change to it. We use a new unsafe func here called 'SliceData'.
	// This will return a pointer to the underlying array of the argument
	// slice. Now that we have a pointer to the array, we can add the size
	// of one string to the pointer to get the address of the second element.
	// This means you could say 2*unsafe.Sizeof("") to get to the third
	// element in this example if that is helpful at all.
	secondAnimal := (*string)(unsafe.Add(
		unsafe.Pointer(unsafe.SliceData(u.animals)), unsafe.Sizeof(""),
	))
	p(u) // {name:bradford age:34 animals:[missy ellie toby]}

	*secondAnimal = "carlos"
	p(u) // {name:bradford age:34 animals:[missy carlos toby]}

	// -------------------------------------------------------------------------

	fruits := []string{"apples", "oranges", "bananas", "kansas"}

	// Get an unsafe.Pointer to the slice data
	start := unsafe.Pointer(unsafe.SliceData(fruits))

	// Get the size of an item in the slice. This could also be
	// written as 'size := unsafe.Sizeof("")' here.
	size := unsafe.Sizeof(fruits[0])

	// Here we loop over the slice and print the data in each item.
	// Arrays in Go are stored contiguously and sequentially in memory,
	// so we are able to directly access each item through indexing:
	//
	// 'base_address + (index * size_of_element)'.
	//
	// In each iteration, we take the pointer to the array data ('start')
	// and add the (index * size_of_an_item) to get the address of each
	// item along the block of memory. Finally, we cast the item to an
	// 'unsafe.Pointer' and then into a '*string' to print it.
	for i := 0; i < len(fruits); i++ {
		p(*(*string)(unsafe.Add(start, uintptr(i)*size)))
	}
	// apples
	// oranges
	// bananas
	// kansas

	// -------------------------------------------------------------------------

	privUser := priv.NewUser()
	p(privUser) // {name:admin age:50 animals:[roger barry melissa]}

	name := (*string)(unsafe.Pointer(&privUser))
	*name = "bradford"
	p(privUser) // {name:bradford age:50 animals:[roger barry melissa]}

	age = (*int)(unsafe.Add(unsafe.Pointer(&privUser), unsafe.Sizeof("")))
	*age = 20
	p(privUser) // {name:bradford age:20 animals:[roger barry melissa]}

	slcPtr := (*[]string)(unsafe.Add(
		unsafe.Pointer(&privUser), (unsafe.Sizeof("") + unsafe.Sizeof(int(0))),
	))
	p(*slcPtr) // [roger barry melissa]

	start = unsafe.Pointer(unsafe.SliceData(*slcPtr))
	size = unsafe.Sizeof("")

	for i := 0; i < len(*slcPtr); i++ {
		p(*(*string)(unsafe.Add(start, uintptr(i)*size)))
	}
	// roger
	// barry
	// melissa

	// -------------------------------------------------------------------------

	// When converting between strings and byte slices in Go, the standard
	// library's string() and []byte{} are commonly used for their safety and
	// simplicity. These methods create a new copy of the data, ensuring that
	// the original data remains immutable and that the type safety is
	// maintained. However, this also means that every conversion involves
	// memory allocation and copying, which can be a performance concern in
	// certain high-efficiency scenarios.

	// StringToByteSlice
	myString := "neato burrito"
	byteSlice := unsafe.Slice(unsafe.StringData(myString), len(myString))
	p(byteSlice) // [110 101 97 116 111 32 98 117 114 114 105 116 111]

	// ByteSliceToString
	myBytes := []byte{
		115, 111, 32, 109, 97, 110, 121, 32, 110,
		101, 97, 116, 32, 98, 121, 116, 101, 115,
	}
	str := unsafe.String(unsafe.SliceData(myBytes), len(myBytes))
	p(str) // so many neat bytes

	// While unsafe provides a high-performance alternative for string-byte
	// slice conversions, it should be used judiciously, primarily when
	// profiling indicates that memory allocations in string conversions
	// are a significant bottleneck. The benefits of using unsafe for
	// these conversions must be weighed against the increased complexity
	// and potential risks.
}

// func main() {
// 	// Allocate integer, get pointer to it:
// 	x := 10
// 	xPtr := &x

// 	// Get a uintptr of the address of x:
// 	xUintPtr := uintptr(unsafe.Pointer(xPtr))

// 	// ---------------------------------------------------------------
// 	// At this point, `x` is unused and so could be garbage collected.
// 	// If that happens, we then have an uintptr (integer) that when
// 	// casted back to an unsafe.Pointer, points to to some invalid
// 	// piece of memory.
// 	// ---------------------------------------------------------------

// 	fmt.Println(*(*int)(unsafe.Pointer(xUintPtr))) // possible misuse of unsafe.Pointer
// }
