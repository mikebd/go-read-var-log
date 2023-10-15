package config

import (
	"flag"
)

type Arguments struct {
	HttpPort         int
	LargeFileBytes   int64
	LogTimestamps    bool
	NumberOfLogLines int
}

func ParseArguments() Arguments {
	args := Arguments{}
	configure(&args)

	flag.IntVar(&args.HttpPort, "p", HttpPort, "HTTP port")
	flag.Int64Var(&args.LargeFileBytes, "l", LargeFileBytes, "Large file bytes")
	flag.BoolVar(&args.LogTimestamps, "lt", false, "Log timestamps (UTC)")
	flag.IntVar(&args.NumberOfLogLines, "n", NumberOfLogLines, "Number of log lines to return")
	flag.Parse()

	return args
}
