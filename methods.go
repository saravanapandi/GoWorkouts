package main

import "fmt"

type shape struct {
	x int
	y int
}

func (s shape) calc() int {
	return 2 * (s.x + s.y)
}

func main() {
	res1 := shape{x: 10, y: 20}
	fmt.Println(res1.calc())

}
