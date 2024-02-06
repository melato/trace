package trace

type funcOption struct {
	prefix string
	name   string
	//  The flags to set if the option is enabled
	fn func(bool)

	description string
}

func newFuncOption(prefix, name string, fn func(bool), description string) Option {
	return &funcOption{prefix, name, fn, description}
}

// Name is the identifier used to enable the option
func (t *funcOption) Name() string {
	if t.prefix != "" {
		if t.name != "" {
			return t.prefix + "." + t.name
		} else {
			return t.prefix
		}
	}
	return t.name
}

// Description is a short description that explains what enabling of this option does.  It is displayed when the options usage is displayed.
func (t *funcOption) Description() string {
	return t.description
}

func (t *funcOption) Enable(on bool) {
	t.fn(on)
}
