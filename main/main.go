package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	defer log.Println("Done")

	args := parseArguments()
	initializeLogging(args.logTimestamps)
	validateUsage(args)

	err := run(args)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(args arguments) error {
	log.Println("Running:", os.Args)

	// TODO: Do something useful here

	return nil
}

func initializeLogging(logTimestamps bool) {
	if logTimestamps {
		log.SetFlags(log.LUTC | log.LstdFlags | log.Lshortfile)
	} else {
		log.SetFlags(log.Lshortfile)
	}
	log.SetOutput(os.Stdout)
}

func validateUsage(args arguments) {
	var invalidUsage bool

	/*
		if len(args.mandatoryStringArgument) == 0 {
			invalidUsage = true
			log.Println("Missing mandatory command line argument: -arg")
		}
	*/

	if invalidUsage {
		flag.Usage()
		os.Exit(2) // Same code used when a flag is provided but not defined
	}
}
