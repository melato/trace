package trace

// Flags - a structure that can be used to setup trace options.
// It is designed to be used as the flags of the top-level command in a CLI application.
// Its public interface does not use any trace data types, so that
// it does not add such dependencies in all modules that you want to trace.
type Flags struct {
	Options
	Trace string `name:"trace" usage:"comma-separated list of trace options.  Use '.' for help"`
}

// nop
func (t *Flags) Init() error {
	return nil
}

// Configured sets the trace flags
func (t *Flags) Configured() error {
	return Set(t.Trace, t.GetOptions()...)
}
