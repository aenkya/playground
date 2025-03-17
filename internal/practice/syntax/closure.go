package syntax

import "fmt"

func closure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		println(pos(i), neg(-2*i))
	}
}

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

func arraySize() {
	x := [3]int{1, 2}
	updateElement(x)
	fmt.Println(x)
}

func updateElement(x [3]int) {
	x[2] = 5
}
