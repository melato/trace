package trace

// Opt is an Option implementation that sets boolean flags.  Use the function T() to construct options conveniently.
type Opt struct {
	name string
	//  The flags to set if the option is enabled
	flags []*bool
	desc  string
}

// Name implements Option.Name()
func (t *Opt) Name() string { return t.name }

// Enable implements Option.Enable().  It sets the flag values to true.
func (t *Opt) Enable(on bool) {
	for _, flag := range t.flags {
		*flag = on
	}
}

// Description implements Option.Description().
func (t *Opt) Description() string { return t.desc }

// Desc specifies an optional description for the option.
func (t *Opt) Desc(description string) *Opt {
	t.desc = description
	return t
}

// T is a convenience method to construct an option.  It can be chained with .Desc() to specify the option description.
func T(name string, flag ...*bool) *Opt {
	return &Opt{name: name, flags: flag}
}
