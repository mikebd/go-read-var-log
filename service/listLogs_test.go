package service

import (
	"fmt"
)

func ExampleListLogs_logDirEmpty() {
	fmt.Println(ListLogs("testdata/logDirEmpty"))
	// Output:
	// [] <nil>
}

func ExampleListLogs_logDir1() {
	fmt.Println(ListLogs("testdata/logDir1"))
	// Output:
	// [10KiB.log] <nil>
}
