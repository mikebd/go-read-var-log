package service

import (
	"fmt"
	"strings"
)

// Log output is in natural order at the service layer.
// Controllers may reverse the order as required by their clients.

func ExampleGetLog_logDir1_10KiB_log() {
	lines, _ := GetLog("testdata/logDir1", "10KiB.log", 2)
	fmt.Println(strings.Join(lines, "\n"))
	// Output:
	// 2023-10-06T15:18:24.408740Z|info |olaret esanus ivo hey enug tewos ebad it u tuge po elora e iwemat o
	// 2023-10-06T15:18:24.408762Z|debug|tucev uho e u ela opif ce igodeto hudegor ivosu ehab eaunopi balohan tagused gicefas
}

func ExampleGetLog_logDir1_1line_log() {
	// Requesting more lines than available returns all available lines.
	lines, _ := GetLog("testdata/logDir1", "1line.log", 10)
	fmt.Println(strings.Join(lines, "\n"))
	// Output:
	// 2023-10-06T15:18:24.406350Z|debug|toyeni vate riwehu ato ped afe ral bo h redi esohet sir moyireh nema lidef
}
