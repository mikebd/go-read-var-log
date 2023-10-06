package config

import (
	"flag"
)

type Arguments struct {
	HttpPort      int
	LogTimestamps bool
}

func ParseArguments() Arguments {
	args := Arguments{}

	flag.IntVar(&args.HttpPort, "p", HttpPort, "HTTP port")
	flag.BoolVar(&args.LogTimestamps, "lt", false, "Log timestamps (UTC)")
	flag.Parse()

	return args
}
