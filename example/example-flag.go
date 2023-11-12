package main

import (
	_ "embed"
	"example/util"
	"flag"

	"melato.org/trace"
)

func main() {
	var app trace.App
	app.AddFuncsDesc("", util.TraceFuncs(), util.TraceDescriptions)
	flag.StringVar(&app.Trace, "trace", "", "comma-separated list of trace options")
	flag.Parse()
	app.Configured()
	util.Run()
}
