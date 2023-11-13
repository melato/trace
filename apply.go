package trace

import (
	"fmt"
	"sort"
	"strings"
)

type optionSorter []Option

func (t optionSorter) Len() int      { return len(t) }
func (t optionSorter) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t optionSorter) Less(i, j int) bool {
	iname := t[i].Name()
	jname := t[j].Name()
	iDot := strings.Index(iname, ".") >= 0
	jDot := strings.Index(jname, ".") >= 0
	if iDot == jDot {
		return iname < jname
	}
	return jDot
}

func print(traces []Option) {
	cp := make([]Option, len(traces))
	copy(cp, traces)
	sort.Sort(optionSorter(cp))
	nameLen := 1
	for _, tr := range cp {
		name := tr.Name()
		w := len(name)
		if w > nameLen {
			nameLen = w
		}
	}
	fmt.Printf("Available traces:\n")
	for _, tr := range cp {
		fmt.Printf(" %-*s %s\n", nameLen, tr.Name(), tr.Description())
	}
	fmt.Printf(" %-*s %s\n", nameLen, ".", "all of the above")
}

// Set parses the trace string, which consists of comma-separated option names.
// It then goes through the provided options and enables the options whose names are included in the list of names.
// If a name does not match an option, the list of available options is printed to stdout.
//
// The special name "." is used to enable all options.
func Set(traceString string, option ...Option) error {
	if traceString == "" {
		return nil
	}
	names := make(map[string]bool)
	for _, name := range strings.Split(traceString, ",") {
		names[name] = true
	}

	all := names["."]
	validNames := map[string]bool{".": true}
	//validNames := make(map[string]bool)
	for _, t := range option {
		name := t.Name()
		if all || names[name] {
			t.Enable(true)
		}
		validNames[name] = true
	}

	// check for invalid specifications
	for name, _ := range names {
		if !validNames[name] {
			print(option)
			return fmt.Errorf("unknown trace: %s", name)
		}
	}
	return nil
}
