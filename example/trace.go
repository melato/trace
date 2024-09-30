package example

import (
	_ "embed"
)

type Trace struct{}

//go:embed trace.yaml
var traceDescriptions []byte

func (t *Trace) Funcs() map[string]func(bool) {
	return map[string]func(bool){
		"a": func(b bool) { TraceA = b },
		"b": func(b bool) { TraceB = b },
	}
}

func (t *Trace) Descriptions() []byte {
	return traceDescriptions
}
