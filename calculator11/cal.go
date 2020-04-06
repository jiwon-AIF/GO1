package calc

import "math/rand"

func Count(x int) int {
	return x
}

func Random(y int) int {
	return y
}

func main() {
	Count(6)
	Random(rand.Intn(45))
}
