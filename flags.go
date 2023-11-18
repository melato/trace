package trace

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// Flags - a structure that can be used to setup trace options.
// It is designed to be used as the flags of the top-level command in a CLI application.
// Its public interface does not use any trace data types, so that
// it does not add such dependencies in all modules that you want to trace.
type Flags struct {
	traceOptions []Option `name:"-"`
	Trace        string   `name:"trace" usage:"comma-separated list of trace options.  Use ? for help"`
}

func (t *Flags) AddVariables(prefix string, m map[string]*bool) {
	for name, v := range m {
		t.traceOptions = append(t.traceOptions,
			newFuncOption(prefix, name, func(b bool) { *v = b }, ""))
	}
}

// AddFuncsDesc adds trace options from a map from trace names a func that enables tracing for that name.
// If prefix is not empty, it is prepended to the name, with a "." as separator.
// If descriptions is not empty, it contains a YAML map from unprefixed names to descriptions.
func (t *Flags) AddFuncsDesc(prefix string, m map[string]func(bool), descriptionsYaml []byte) {
	var descriptions map[string]string
	if len(descriptionsYaml) > 0 {
		err := yaml.Unmarshal(descriptionsYaml, &descriptions)
		if err != nil {
			fmt.Printf("trace prefix %s descriptions error: %v\n", prefix, err)
		}
	}
	for name, fn := range m {
		t.traceOptions = append(t.traceOptions, newFuncOption(prefix, name, fn, descriptions[name]))
	}
}

func (t *Flags) AddOptions(options ...Option) {
	t.traceOptions = append(t.traceOptions, options...)
}

// AddOptionFuncs adds trace options from a map from trace names a func that enables tracing for that name.
// If prefix is not empty, it is prepended to the name, with a "." as separator.
func (t *Flags) AddFuncs(prefix string, m map[string]func(bool)) {
	t.AddFuncsDesc(prefix, m, nil)
}

func (t *Flags) Init() error {
	return nil
}

func (t *Flags) Configured() error {
	return Set(t.Trace, t.traceOptions...)
}
