package advanced

import "fmt"

func Maps() {
	mp := map[string]int{"a": 1}
	fmt.Println("first map =>", mp)

	mp2 := make(map[int]string)
	mp2[2] = "hello world"
	fmt.Println("map 2 =>", mp2)

	mp3 := map[string][]int{"key1": {1, 2, 3}}
	mp3["key2"] = []int{4, 5, 6}
	fmt.Println("map 3 before override =>", mp3)
	mp3["key2"] = []int{10, 11, 12}
	fmt.Println("map 3 after override =>", mp3)
	delete(mp3, "key2")
	fmt.Println("map 3 after delete prop =>", mp3)

	value, ok := mp["key1"]
	fmt.Println("value is present =>", ok, " ||||   the value is =>", value)

	mp4 := map[uint]uint{}
	n := 100

	for i := range n {
		for d := 1; d <= 5; d++ {
			if i%d == 0 {
				mp4[uint(d)]++
			}
		}
	}

	fmt.Println("map with dividers =>", mp4)
}
