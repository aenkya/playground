package syntax

import (
	"fmt"
	"runtime"
)

type myType int

func Types() {
	var x myType = 8

	fmt.Println(runtime.NumCPU())

	arr := [...]int{3, 5, 2}
	y := arr[0]
	fmt.Println(x, len(arr), y)

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println(sample)

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x", sample[i]) // hexa decimal representation
	}

	fmt.Println()

	for _, rune := range sample {
		fmt.Printf("% d\n", rune)
	}
}
