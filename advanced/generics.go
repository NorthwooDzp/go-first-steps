package advanced

import "fmt"

type number interface{ int | uint | float64 | float32 } // union type

type GenericSlice[T any] []T

type GenericStruct[T any] struct {
	name     string
	metadata T
}

func (g GenericSlice[T]) Print() {
	for index, value := range g {
		fmt.Println("Value at the index", index, "is =>", value)
	}
}

func addNums[T number](x, y T) T {
	return x + y
}

func logger[T any](x T) {
	fmt.Println(x)
}

func getValues[K comparable, V any](mp map[K]V) (values []V) {
	for _, value := range mp {
		values = append(values, value)
	}

	return
}

func Generics() {
	logger(addNums(10, 12))
	logger("Test string")
	mp1 := map[string]string{
		"test":   "test01",
		"notest": "test02",
	}

	logger(getValues(mp1))

	slice1 := GenericSlice[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1.Print()

	struct1 := GenericStruct[string]{name: "Some struct", metadata: "some metadata, can be any"}
	fmt.Println("Struct with generic metadata =>", struct1)
	println("Do we need fmt for that?", "I guess no...")
}
