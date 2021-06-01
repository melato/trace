// Simple facility for initializing trace options
// A trace option typically enables trace output, for debugging, testing, or diagnostics,
// without otherwise changing a program's behaviour
// Trace options are initialized by passing a comma-separated list of option names and option definitions to Set()
// Trace options are conveniently specified by calling the T() function,
// but can also be specified by implementing the Option interface.
package trace

import (
	"errors"
	"fmt"
	"strings"
)

// Option specifies a trace option.
type Option interface {
	// Name is the identifier used to enable the option
	Name() string
	// Description is a short description that explains what enabling of this option does.  It is displayed when the options usage is displayed.
	Description() string
	// Enable this option
	Enable()
}

// Opt is an Option that sets boolelan flags.  Use the function T() to construct options conveniently.
type Opt struct {
	name string
	//  The flags to set if the option is enabled
	flags []*bool
	desc  string
}

// Name implements Option.Name()
func (t *Opt) Name() string { return t.name }

// Enable implements Option.Enable().  It sets the flag values to true.
func (t *Opt) Enable() {
	for _, flag := range t.flags {
		*flag = true
	}
}

// Description implements Option.Description().
func (t *Opt) Description() string { return t.desc }

// Desc specifies an optional description for the option.  It shows up in the list of options.
func (t *Opt) Desc(description string) *Opt {
	t.desc = description
	return t
}

// T is a convenience method to construct an option.  It can be chained with .Desc() to specify the option description.
func T(name string, flag ...*bool) *Opt {
	return &Opt{name: name, flags: flag}
}

// Flags is a convenience type that can be incorporated in flag structs to define a -t flag for trace
type Flags struct {
	Trace string `name:"trace" usage:"comma-separated trace options (. = all)"`
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
			t.Enable()
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
