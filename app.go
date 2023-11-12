package trace

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// App - a structure that can be used to setup
// It is designed to be used as the flags of the top-level command in a CLI application.
type App struct {
	TraceOptions []Option `name:"-"`
	Flags
}

func (t *App) AddOptionVariables(prefix string, m map[string]*bool) {
	for name, v := range m {
		if prefix != "" {
			name = prefix + "." + name
		}
		t.TraceOptions = append(t.TraceOptions,
			T(name, v))
	}
}

// AddOptionFuncs adds trace options from a map from trace names a func that enables tracing for that name.
// If prefix is not empty, it is prepended to the name, with a "." as separator.
// If descriptions is not empty, it contains a YAML map from unprefixed names to descriptions.
func (t *App) AddOptionFuncsDesc(prefix string, m map[string]func(bool), descriptionsYaml []byte) {
	var descriptions map[string]string
	if len(descriptionsYaml) > 0 {
		err := yaml.Unmarshal(descriptionsYaml, &descriptions)
		if err != nil {
			fmt.Printf("trace prefix %s descriptions error: %v\n", prefix, err)
		}
	}
	for name, fn := range m {
		var qname = name
		if prefix != "" {
			qname = prefix + "." + name
		}
		t.TraceOptions = append(t.TraceOptions, &funcOption{qname, fn, descriptions[name]})
	}
}

// AddOptionFuncs adds trace options from a map from trace names a func that enables tracing for that name.
// If prefix is not empty, it is prepended to the name, with a "." as separator.
func (t *App) AddOptionFuncs(prefix string, m map[string]func(bool)) {
	t.AddOptionFuncsDesc(prefix, m, nil)
}

type funcOption struct {
	name string
	//  The flags to set if the option is enabled
	fn func(bool)

	description string
}

// Name is the identifier used to enable the option
func (t *funcOption) Name() string {
	return t.name
}

// Description is a short description that explains what enabling of this option does.  It is displayed when the options usage is displayed.
func (t *funcOption) Description() string {
	return t.description
}

func (t *funcOption) Enable(on bool) {
	t.fn(on)
}

func (t *App) Init() error {
	return nil
}

func (t *App) Configured() error {
	return t.Flags.Set(t.TraceOptions...)
}
