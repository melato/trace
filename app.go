package trace

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
func (t *App) AddOptionFuncs(prefix string, m map[string]func(bool)) {
	for name, fn := range m {
		if prefix != "" {
			name = prefix + "." + name
		}
		t.TraceOptions = append(t.TraceOptions, &funcOption{name, fn})
	}
}

type funcOption struct {
	name string
	//  The flags to set if the option is enabled
	fn func(bool)
}

// Name is the identifier used to enable the option
func (t *funcOption) Name() string {
	return t.name
}

// Description is a short description that explains what enabling of this option does.  It is displayed when the options usage is displayed.
func (t *funcOption) Description() string {
	return ""
}

func (t *funcOption) Enable() {
	t.fn(true)
}

func (t *App) Init() error {
	return nil
}

func (t *App) Configured() error {
	return t.Flags.Set(t.TraceOptions...)
}
