
A very simple mechanism for controlling tracing in Go is to use global bool
variables:

```
var Trace bool

func foo() {
	if Trace {
		fmt.Printf("foo\n")
	}
}
```

# initialization
This package just provides a mechanism to initialize such Trace variables:

```
	var flags trace.Flags
	flags.Trace = "a"  // use a command-line option to set this
	
	err := t.Flags.Set(
		trace.T("a", &TraceA),
		trace.T("b", &TraceB),
	)
	// sets TraceA to true
```

The trace package has no dependencies to other packages,
but the trace.Flags struct can be used with the melato.org/command package
(via reflection), as shown in example.go.


