package service

import (
	"fmt"
)

// based on config.MaxResultLines = 2_000

func Example_maxLines_none() {
	fmt.Println(maxLines(&GetLogParams{}))
	// Output:
	// 2000
}

func Example_maxLines_zero() {
	fmt.Println(maxLines(&GetLogParams{MaxLines: 0}))
	// Output:
	// 2000
}

func Example_maxLines_negative() {
	fmt.Println(maxLines(&GetLogParams{MaxLines: -1}))
	// Output:
	// 2000
}

func Example_maxLines_one() {
	fmt.Println(maxLines(&GetLogParams{MaxLines: 1}))
	// Output:
	// 1
}

func Example_maxLines_two() {
	fmt.Println(maxLines(&GetLogParams{MaxLines: 2}))
	// Output:
	// 2
}

func Example_maxLines_nineteen_ninety_nine() {
	fmt.Println(maxLines(&GetLogParams{MaxLines: 1_999}))
	// Output:
	// 1999
}

func Example_maxLines_two_thousand() {
	fmt.Println(maxLines(&GetLogParams{MaxLines: 2_000}))
	// Output:
	// 2000
}

func Example_maxLines_two_thousand_one() {
	fmt.Println(maxLines(&GetLogParams{MaxLines: 2_001}))
	// Output:
	// 2000
}
