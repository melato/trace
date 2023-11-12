package trace

// Option specifies something that can be turned on to enable debugging output.
type Option interface {
	// Name is the identifier used to enable the option
	Name() string
	// Description is a short description that explains what enabling of this option does.  It is displayed when the options usage is displayed.
	Description() string
	// Enable or disable this option
	Enable(on bool)
}
