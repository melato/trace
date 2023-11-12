package trace

// Flags is a convenience type that can be incorporated in flag structs to define a -t flag for trace
type Flags struct {
	Trace string `name:"trace" usage:"comma-separated list of trace options"`
}

// Set is short for Set(Flags.Trace, traces)
func (t *Flags) Set(traces ...Option) error {
	return Set(t.Trace, traces...)
}
