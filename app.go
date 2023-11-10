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

func (t *App) Init() error {
	return nil
}

func (t *App) Configured() error {
	return t.Flags.Set(t.TraceOptions...)
}
