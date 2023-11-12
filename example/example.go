package main

import (
	_ "embed"
	"example/util"

	"melato.org/command"
	"melato.org/trace"
)

func main() {
	cmd := &command.SimpleCommand{}
	var app trace.App
	app.AddFuncsDesc("", util.TraceFuncs(), util.TraceDescriptions)
	cmd.Flags(&app)
	cmd.Command("run").RunFunc(util.Run)
	command.Main(cmd)
}
