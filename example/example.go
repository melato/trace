package main

import (
	_ "embed"
	"example/util"

	"melato.org/command"
	"melato.org/trace"
)

func main() {
	cmd := &command.SimpleCommand{}
	var flags trace.Flags
	flags.AddFuncsDesc("", util.TraceFuncs(), util.TraceDescriptions)
	cmd.Flags(&flags)
	cmd.Command("run").RunFunc(util.Run)
	command.Main(cmd)
}
