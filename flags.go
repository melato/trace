package trace

import (
	"errors"
	"fmt"
	"strings"
)

// Flags is a convenience type that can be incorporated in flag structs to define a -t flag for trace
type Flags struct {
	Trace string `name:"trace" usage:"comma-separated list of trace options"`
}

// Set is short for Set(Flags.Trace, traces)
func (t *Flags) Set(traces ...Option) error {
	return Set(t.Trace, traces...)
}

func print(traces []Option) {
	fmt.Printf("Available traces:\n")
	nameLen := 1
	for _, t := range traces {
		w := len(t.Name())
		if w > nameLen {
			nameLen = w
		}
	}
	for _, t := range traces {
		fmt.Printf(" %-*s %s\n", nameLen, t.Name(), t.Description())
	}
	fmt.Printf(" %-*s %s\n", nameLen, ".", "all of the above")
}

// Set parses the trace string, which consists of comma-separated option names.
// It then goes through the provided options and enables the options whose names are included in the list of names.
// If a name does not match an option, the list of available options is printed to stdout.
//
// The special name "." is used to enable all options.
func Set(traceString string, option ...Option) error {
	if traceString == "" {
		return nil
	}
	names := make(map[string]bool)
	for _, name := range strings.Split(traceString, ",") {
		names[name] = true
	}

	all := names["."]
	validNames := map[string]bool{".": true}
	//validNames := make(map[string]bool)
	for _, t := range option {
		name := t.Name()
		if all || names[name] {
			t.Enable(true)
		}
		validNames[name] = true
	}

	// check for invalid specifications
	for name, _ := range names {
		if !validNames[name] {
			msg := "unknown trace: " + name
			//fmt.Println(msg)
			print(option)
			return errors.New(msg)
		}
	}
	return nil
}
