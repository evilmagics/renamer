package main

import (
	"flag"
)

type Args struct {
	Dir    string
	Prefix string
	Suffix string
}

func parseArgs() Args {
	args := new(Args)
	flag.StringVar(&args.Dir, "dir", ".", "Directory to search")
	flag.StringVar(&args.Prefix, "prefix", "", "Rename files with prefix specified")
	flag.StringVar(&args.Suffix, "suffix", "", "Rename files with suffix specified")

	flag.Parse()

	return *args
}
