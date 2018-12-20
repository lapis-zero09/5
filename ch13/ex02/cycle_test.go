package cycle

import (
	"testing"
)

func TestIsCycle(t *testing.T) {

	type P *P
	var p P
	p = &p

	type S []S
	s := make(S, 1)
	s[0] = s

	type Cycle struct {
		Tail *Cycle
	}
	var c Cycle
	c = Cycle{&c}

	ts := []struct {
		i        interface{}
		expected bool
	}{
		{*new(int), false},
		{p, true},

		{[]int{}, false},
		{s, true},

		{Cycle{}, false},
		{c, true},
	}

	for _, tc := range ts {
		got := IsCycle(tc.i)
		if got != tc.expected {
			t.Errorf("unexpected result. expected: %v, but got: %v", tc.expected, got)
		}
	}
}
