// example-flag uses the flag package for command-line arguments
package main

import (
	_ "embed"
	"example"
	"flag"

	"melato.org/trace"
)

func main() {
	var flags trace.Flags
	flags.Add("", &example.Trace{})
	flag.StringVar(&flags.Trace, "trace", "", "comma-separated list of trace options")
	flag.Parse()
	flags.Configured()
	example.Run()
}
