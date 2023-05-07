package main

import "fmt"

func main() {
	mapData := make(map[string]int)
	mapData["a"] = 1
	mapData["b"] = 2
	fmt.Println(mapData)

	delete(mapData, "b")
	fmt.Println(mapData)

	sliceData := make([]string, 3)
	sliceData[0] = "a"
	sliceData[1] = "b"
	sliceData[2] = "c"
	fmt.Println(sliceData)

	l := sliceData[:2]
	fmt.Println(l)

	j := sliceData[2:]
	fmt.Println(j)
}
