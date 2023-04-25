package main

import (
	"fmt"

	"melato.org/command"
	"melato.org/trace"
)

var TraceA bool
var TraceB bool

type App struct {
	trace.Flags
}

func (t *App) Init() error {
	return nil
}

func (t *App) Configured() error {
	return t.Flags.Set(
		trace.T("a", &TraceA),
		trace.T("b", &TraceB),
	)
}

func (t *App) Run() error {
	if TraceA {
		fmt.Printf("a\n")
	}
	if TraceB {
		fmt.Printf("b\n")
	}
	return nil
}

func main() {
	cmd := &command.SimpleCommand{}
	var app App
	cmd.Flags(&app).RunFunc(app.Run)

	command.Main(cmd)
}
