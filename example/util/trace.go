package util

import (
	_ "embed"
)

//go:embed trace.yaml
var TraceDescriptions []byte

func TraceFuncs() map[string]func(bool) {
	return map[string]func(bool){
		"a": func(b bool) { TraceA = b },
		"b": func(b bool) { TraceB = b },
	}
}
