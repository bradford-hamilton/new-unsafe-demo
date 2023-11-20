package priv

type s struct {
	foo string
	bar int
	baz []int
}

func NewS() s {
	return s{
		foo: "bar",
		bar: 1337,
		baz: []int{100, 150, 200, 250},
	}
}
