package trace

import (
	"fmt"
	"regexp"
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
}

func SetOptions(names []string, option []Option) error {
	if len(names) == 1 && names[0] == "." {
		print(option)
		fmt.Printf("Use %% for wildcard\n")
		return fmt.Errorf("(exit)")
	}
	optionNames := make(map[string]bool)
	for _, t := range option {
		optionNames[t.Name()] = true
	}
	var patterns []*regexp.Regexp
	nameMap := make(map[string]bool) // non regexp
	for _, name := range names {
		if strings.ContainsAny(name, "*%") {
			expr := strings.ReplaceAll(name, ".", "\\.")
			expr = strings.ReplaceAll(expr, "*", ".*")
			expr = strings.ReplaceAll(expr, "%", ".*")
			expr = "^" + expr + "$"
			re, err := regexp.Compile(expr)
			if err != nil {
				return fmt.Errorf("expr: %s: %w", expr, err)
			}
			patterns = append(patterns, re)
		} else {
			if !optionNames[name] {
				fmt.Printf("use '.' for list of trace names\n")
				return fmt.Errorf("unknown trace: %s", name)
			}
			nameMap[name] = true
		}
	}

	for _, t := range option {
		name := t.Name()
		var enable bool
		if nameMap[name] {
			enable = true
		} else {
			for _, re := range patterns {
				if re.MatchString(name) {
					enable = true
					break
				}
			}
		}
		if enable {
			t.Enable(true)
			fmt.Printf("enable trace: %s\n", name)
		}
	}
	return nil

}

// Set parses the trace string, which consists of comma-separated option names.
// It then goes through the provided options and enables the options whose names are included in the list of names.
// If a name does not match an option, the list of available options is printed to stdout.
//
// The special name "%" is used to enable all options.
func Set(traceString string, option ...Option) error {
	if traceString == "" {
		return nil
	}
	names := strings.Split(traceString, ",")
	return SetOptions(names, option)
}
