package trace

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// Options - contains available trace options,
// and methods to add options.
// Its public interface does not use any trace data types, so that
// it does not add such dependencies in all modules that you want to trace.
type Options struct {
	traceOptions []Option `name:"-"`
}

func (t *Options) GetOptions() []Option {
	return t.traceOptions
}

func (t *Options) AddVariables(prefix string, m map[string]*bool) {
	for name, v := range m {
		t.traceOptions = append(t.traceOptions,
			newFuncOption(prefix, name, func(b bool) { *v = b }, ""))
	}
}

// AddFuncsDesc adds trace options from a map from trace names a func that enables tracing for that name.
// If prefix is not empty, it is prepended to the name, with a "." as separator.
// If descriptions is not empty, it contains a YAML map from unprefixed names to descriptions.
func (t *Options) AddFuncsDesc(prefix string, m map[string]func(bool), descriptionsYaml []byte) {
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

func (t *Options) AddOptions(options ...Option) {
	t.traceOptions = append(t.traceOptions, options...)
}

// AddOptionFuncs adds trace options from a map from trace names a func that enables tracing for that name.
// If prefix is not empty, it is prepended to the name, with a "." as separator.
func (t *Options) AddFuncs(prefix string, m map[string]func(bool)) {
	t.AddFuncsDesc(prefix, m, nil)
}

func (t *Options) Add(prefix string, v any) {
	switch x := v.(type) {
	case Funcs:
		var descriptions []byte
		desc, hasDesc := x.(Descriptions)
		if hasDesc {
			descriptions = desc.Descriptions()

		}
		t.AddFuncsDesc(prefix, x.Funcs(), descriptions)
	case map[string]*bool:
		t.AddVariables(prefix, x)
	default:
		fmt.Printf("trace Add(%s): unsupported type: %T\n", prefix, v)
	}
}

func (t *Options) Set(names []string) error {
	if len(names) == 0 {
		return nil
	}
	return SetOptions(names, t.GetOptions())
}

func (t *Options) SetString(names string) error {
	if len(names) == 0 {
		return nil
	}
	return Set(names, t.GetOptions()...)
}
