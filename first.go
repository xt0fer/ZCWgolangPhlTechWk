package main //every go file has a package clause

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, 世界")
}

func add(x int, y int) int {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
