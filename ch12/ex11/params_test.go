package params

import "testing"

func TestPack(t *testing.T) {
	ts := []struct {
		i        interface{}
		expected string
	}{
		{
			i: &struct {
				aaa string
			}{
				"bbb",
			},
			expected: "aaa=bbb",
		},
		{
			i: &struct {
				b bool
			}{
				true,
			},
			expected: "b=true",
		},
		{
			i: &struct {
				aaa string
				bbb string
			}{
				"bbb",
				"",
			},
			expected: "aaa=bbb&bbb=",
		},
		{
			i: &struct {
				v []string
			}{
				[]string{"aaa", "bbb", "ccc"},
			},
			expected: "v=aaa&v=bbb&v=ccc",
		},
		{
			i: &struct {
				num int
			}{
				123,
			},
			expected: "num=123",
		},
	}

	for _, tc := range ts {
		if got := Pack(tc.i); got != tc.expected {
			t.Errorf("unexpected query string. expected: %v, but got: %v", tc.expected, got)
		}
	}
}
