package service

import (
	"fmt"
)

func ExampleListLogs_logDir1() {
	fmt.Println(ListLogs("testdata/logDir1"))
	// Output:
	// [10KiB.log 1line.log] <nil>
}
