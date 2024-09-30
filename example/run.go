package example

import (
	"fmt"
)

var TraceA bool
var TraceB bool

func Run() error {
	if TraceA {
		fmt.Printf("a\n")
	}
	if TraceB {
		fmt.Printf("b\n")
	}
	return nil
}
