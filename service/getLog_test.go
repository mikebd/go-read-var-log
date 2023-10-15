package service

import (
	"fmt"
	"regexp"
	"strings"
)

// Log output is in natural order at the service layer.
// Controllers may reverse the order as required by their clients.

func ExampleGetLog_logDir1_10KiB_log() {
	getLogResult := GetLog(&GetLogParams{
		DirectoryPath: "testdata/logDir1",
		Filename:      "10KiB.log",
		MaxLines:      2,
	})
	fmt.Println(getLogResult.Strategy)
	fmt.Println(strings.Join(getLogResult.LogLines, "\n"))
	// Output:
	// small
	// 2023-10-06T15:18:24.408740Z|info |olaret esanus ivo hey enug tewos ebad it u tuge po elora e iwemat o
	// 2023-10-06T15:18:24.408762Z|debug|tucev uho e u ela opif ce igodeto hudegor ivosu ehab eaunopi balohan tagused gicefas
}

func ExampleGetLog_logDir1_99lines_log_text_filter() {
	getLogResult := GetLog(&GetLogParams{
		DirectoryPath: "testdata/logDir1",
		Filename:      "99lines.log",
		TextMatch:     "9 ",
	})
	fmt.Println(getLogResult.Strategy)
	fmt.Println(strings.Join(getLogResult.LogLines, "\n"))
	// Output:
	// small
	// line 9 yz
	// line 19 yz
	// line 29 yz
	// line 39 yz
	// line 49 yz
	// line 59 yz
	// line 69 yz
	// line 79 yz
	// line 89 yz
	// line 99 yz
}

func ExampleGetLog_logDir1_99lines_log_regex_filter() {
	regex := regexp.MustCompile("[9]\\s.z")
	// Requesting 0 lines returns all available lines.
	getLogResult := GetLog(&GetLogParams{
		DirectoryPath: "testdata/logDir1",
		Filename:      "99lines.log",
		Regex:         regex,
	})
	fmt.Println(getLogResult.Strategy)
	fmt.Println(strings.Join(getLogResult.LogLines, "\n"))
	// Output:
	// small
	// line 9 yz
	// line 19 yz
	// line 29 yz
	// line 39 yz
	// line 49 yz
	// line 59 yz
	// line 69 yz
	// line 79 yz
	// line 89 yz
	// line 99 yz
}

func ExampleGetLog_logDir1_99lines_log_text_and_regex_filter() {
	regex := regexp.MustCompile("[9]\\s.z")
	// Requesting 0 lines returns all available lines.
	getLogResult := GetLog(&GetLogParams{
		DirectoryPath: "testdata/logDir1",
		Filename:      "99lines.log",
		TextMatch:     "7",
		Regex:         regex,
	})
	fmt.Println(getLogResult.Strategy)
	fmt.Println(strings.Join(getLogResult.LogLines, "\n"))
	// Output:
	// small
	// line 79 yz
}

func ExampleGetLog_logDir1_1line_log() {
	// Requesting more lines than available returns all available lines.
	getLogResult := GetLog(&GetLogParams{
		DirectoryPath: "testdata/logDir1",
		Filename:      "1line.log",
		MaxLines:      10,
	})
	fmt.Println(getLogResult.Strategy)
	fmt.Println(strings.Join(getLogResult.LogLines, "\n"))
	// Output:
	// small
	// 2023-10-06T15:18:24.406350Z|debug|toyeni vate riwehu ato ped afe ral bo h redi esohet sir moyireh nema lidef
}
