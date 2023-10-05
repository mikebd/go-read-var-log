package config

import (
	"flag"
)

type Arguments struct {
	LogTimestamps bool
}

func ParseArguments() Arguments {
	args := Arguments{}

	flag.BoolVar(&args.LogTimestamps, "lt", false, "Log timestamps (UTC)")
	flag.Parse()

	return args
}
