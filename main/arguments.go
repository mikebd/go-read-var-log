package main

import (
	"flag"
)

type arguments struct {
	logTimestamps bool
}

func parseArguments() arguments {
	args := arguments{}

	flag.BoolVar(&args.logTimestamps, "lt", false, "Log timestamps (UTC)")
	flag.Parse()

	return args
}
