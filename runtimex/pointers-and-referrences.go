package runtimex

type Book struct {
	title string
	id    int
}

func (b *Book) setTitle(title string) {
	b.title = title
}

func getValuesFromPointers(pointerSlice *[]*int) {
	values := *pointerSlice

	for i, val := range values {
		println("The value at the index", i, "is =>", val)
	}
}

func PointersAndReferrences() {
	changeNum := func(x *int, val int) {
		*x = val
	}

	a := 10
	changeNum(&a, 17)
	println("New a is =>", a)

	b := &a
	c := &b

	changeNum(*c, 34) // **c = 34 will do the same
	println("New a is =>", a, "| b is =>", b, "| c is =>", c)

	book1 := Book{id: 1, title: "Old book"}
	println("Old Book title =>", book1.title)

	book1.setTitle("New book")
	println("New Book title =>", book1.title)

	v1 := 24
	v2 := 32

	values := &[]*int{&v1, &v2}
	getValuesFromPointers(values)
}
