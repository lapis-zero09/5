package intset

import (
	"math/rand"
	"testing"
	"time"
)

func newIntSet(a ...int) *IntSet {
	intSet := &IntSet{}
	for _, x := range a {
		intSet.Add(x)
	}
	return intSet
}

func TestIntSet_Has(t *testing.T) {
	ts := []struct {
		intSet   *IntSet
		expected map[int]bool
	}{
		{
			intSet: newIntSet(1, 10, 100, 1000, 10000),
			expected: map[int]bool{
				0:     false,
				1:     true,
				2:     false,
				9:     false,
				10:    true,
				11:    false,
				99:    false,
				100:   true,
				101:   false,
				999:   false,
				1000:  true,
				1001:  false,
				9999:  false,
				10000: true,
				10001: false,
			},
		},
	}
	for _, tc := range ts {
		for x, expected := range tc.expected {
			if got := tc.intSet.Has(x); got != expected {
				t.Errorf("unexpected Has(%d). expected: %v, but got: %v", x, expected, got)
			}
		}
	}
}

func TestIntSet_Add(t *testing.T) {
	ts := []struct {
		intSet   *IntSet
		adds     []int
		expected map[int]bool
	}{
		{
			intSet: newIntSet(1, 10, 100, 1000),
			adds:   []int{0, 50, 10000},
			expected: map[int]bool{
				0:     true,
				1:     true,
				2:     false,
				9:     false,
				10:    true,
				11:    false,
				49:    false,
				50:    true,
				51:    false,
				99:    false,
				100:   true,
				101:   false,
				999:   false,
				1000:  true,
				1001:  false,
				9999:  false,
				10000: true,
				10001: false,
			},
		},
	}

	for _, tc := range ts {
		for _, x := range tc.adds {
			tc.intSet.Add(x)
		}
		for x, expected := range tc.expected {
			if got := tc.intSet.Has(x); got != expected {
				t.Errorf("unexpected Has(%d). expected: %v, but got: %v", x, expected, got)
			}
		}
	}
}

func TestIntSet_UnionWith(t *testing.T) {
	ts := []struct {
		right, left *IntSet
		expected    map[int]bool
	}{
		{
			right: newIntSet(1, 10, 100, 1000),
			left:  newIntSet(10000, 50000, 100000),
			expected: map[int]bool{
				0:      false,
				1:      true,
				2:      false,
				9:      false,
				10:     true,
				11:     false,
				99:     false,
				100:    true,
				101:    false,
				999:    false,
				1000:   true,
				1001:   false,
				9999:   false,
				10000:  true,
				10001:  false,
				49999:  false,
				50000:  true,
				50001:  false,
				99999:  false,
				100000: true,
				100001: false,
			},
		},
		{
			right: newIntSet(1, 10, 100, 1000),
			left:  newIntSet(2, 10, 1000, 10000),
			expected: map[int]bool{
				0:     false,
				1:     true,
				2:     true,
				3:     false,
				9:     false,
				10:    true,
				11:    false,
				99:    false,
				100:   true,
				101:   false,
				999:   false,
				1000:  true,
				1001:  false,
				9999:  false,
				10000: true,
				10001: false,
			},
		},
		{
			right: newIntSet(1, 10, 100, 1000),
			left:  newIntSet(0, 100),
			expected: map[int]bool{
				0:    true,
				1:    true,
				2:    false,
				9:    false,
				10:   true,
				11:   false,
				99:   false,
				100:  true,
				101:  false,
				999:  false,
				1000: true,
				1001: false,
			},
		},
	}

	for _, tc := range ts {
		tc.left.UnionWith(tc.right)
		for x, expected := range tc.expected {
			if got := tc.left.Has(x); got != expected {
				t.Errorf("unexpected Has(%d). expected: %v, but got: %v", x, expected, got)
			}
		}
	}
}

func TestIntSet_String(t *testing.T) {
	ts := []struct {
		intSet   *IntSet
		expected string
	}{
		{
			intSet:   newIntSet(),
			expected: "{}",
		},
		{
			intSet:   newIntSet(10),
			expected: "{10}",
		},
		{
			intSet:   newIntSet(1, 2, 4, 8, 16, 32),
			expected: "{1 2 4 8 16 32}",
		},
	}

	for _, tc := range ts {
		if got := tc.intSet.String(); got != tc.expected {
			t.Errorf("unexpected string. expected: %q, but got: %q", tc.expected, got)
		}
	}
}

func benchIntSetAdd(b *testing.B, size int) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < b.N; i++ {
		intSet := &IntSet{}
		for j := 0; j < size; j++ {
			intSet.Add(rand.Intn(size))
		}
	}
}

func BenchmarkIntSetAdd1(b *testing.B) {
	benchIntSetAdd(b, 1)
}

func BenchmarkIntSetAdd10(b *testing.B) {
	benchIntSetAdd(b, 10)
}

func BenchmarkIntSetAdd100(b *testing.B) {
	benchIntSetAdd(b, 100)
}

func BenchmarkIntSetAdd1000(b *testing.B) {
	benchIntSetAdd(b, 1000)
}

func BenchmarkIntSetAdd10000(b *testing.B) {
	benchIntSetAdd(b, 10000)
}