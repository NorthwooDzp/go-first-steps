package advanced

import "fmt"

type Shape interface {
	getPerimeter() uint
}

type Triangle struct {
	a, b, c uint
}

type Square struct {
	a uint
}

func (t Triangle) getPerimeter() uint {
	return t.a + t.b + t.c
}

func (s Square) getPerimeter() uint {
	return 4 * s.a
}

func Interfaces() {
	var sh1 Shape = Triangle{3, 4, 5}
	fmt.Println("Shape 1 =>", sh1)
	fmt.Println("Perimeter of the sh1 =>", sh1.getPerimeter())
	// sh1.a, sh1.b and sh1.c are not available properties, they are incapsulated

	var sh2 Shape = Square{5}
	fmt.Println("Shape 2 =>", sh2)
	fmt.Println("Perimeter of the sh2 =>", sh2.getPerimeter())

	var shapes []Shape = []Shape{sh1, sh2}
	fmt.Println("Shapes in one slice =>", shapes) // allow to add in the same slice different by type structs
	perimeters := uint(0)

	for _, shape := range shapes {
		perimeters += shape.getPerimeter()
	}

	fmt.Println("Total perimeter of shapes is =>", perimeters)
}
