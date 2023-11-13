package trace

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	options := []Option{
		T("a.b"),
		T("b"),
	}
	sort.Sort(optionSorter(options))
	if options[0].Name() != "b" {
		for _, opt := range options {
			fmt.Printf("%s\n", opt.Name())
		}
		t.Fail()
	}
}
