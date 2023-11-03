package service

import (
	"fmt"
)

func Example_reduceToMaxLines_0() {
	fmt.Println([]string{})
	fmt.Println(reduceToMaxLines([]string{"a", "b", "c"}, 0))
	// Output:
	// []
	// []
}

func Example_reduceToMaxLines_1() {
	fmt.Println(reduceToMaxLines([]string{"a", "b", "c"}, 1))
	// Output:
	// [c]
}

func Example_reduceToMaxLines_2() {
	fmt.Println(reduceToMaxLines([]string{"a", "b", "c"}, 2))
	// Output:
	// [b c]
}

func Example_reduceToMaxLines_3() {
	fmt.Println(reduceToMaxLines([]string{"a", "b", "c"}, 3))
	// Output:
	// [a b c]
}
