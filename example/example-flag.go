package main

import (
	_ "embed"
	"example/util"
	"flag"

	"melato.org/trace"
)

func main() {
	var flags trace.Flags
	flags.AddFuncsDesc("", util.TraceFuncs(), util.TraceDescriptions)
	flag.StringVar(&flags.Trace, "trace", "", "comma-separated list of trace options")
	flag.Parse()
	flags.Configured()
	util.Run()
}
