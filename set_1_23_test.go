//go:build go1.23
// +build go1.23

package mapset

import (
	"testing"
)

func Test_Elements123(t *testing.T) {
	a := NewSet("Z", "Y", "X", "W")
	b := NewSet[string]()
	for elem := range Elements(a) {
		b.Add(elem)
	}

	if !a.Equal(b) {
		t.Error("The sets are not equal after iterating (Each) through the first set")
	}

	var count int
	for range Elements(a) {
		if count == 2 {
			break
		}
		count++
	}

	if count != 2 {
		t.Error("Iteration should stop on the way")
	}
}
