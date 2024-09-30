// example uses the melato.org/command package for command-line arguments
package main

import (
	_ "embed"

	"example"

	"melato.org/command"
	"melato.org/trace"
)

func main() {
	cmd := &command.SimpleCommand{}
	var flags trace.Flags
	flags.Add("", &example.Trace{})
	cmd.Flags(&flags)
	cmd.Command("run").RunFunc(example.Run)
	command.Main(cmd)
}
