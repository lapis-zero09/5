package sexpr

import (
	"reflect"
	"testing"
)

func TestMarshalTagsexpr(t *testing.T) {
	type testT struct {
		A string
		B string `sexpr:"bTag"`
	}
	ts := []struct {
		i        testT
		expected string
	}{
		{
			i: testT{
				A: "a",
				B: "bar",
			},
			expected: `((A "a") (bTag "bar"))`,
		},
		{
			i: testT{
				B: "bar",
			},
			expected: `( (bTag "bar"))`,
		},
		{
			i: testT{
				A: "bar",
			},
			expected: `((A "bar"))`,
		},
		{
			i:        testT{},
			expected: `()`,
		},
	}

	for _, tc := range ts {
		b, err := Marshal(tc.i)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if got != tc.expected {
			t.Errorf("unexpected s expr. expected: %q, but got: %q", tc.expected, got)
		}
	}
}

func TestUnmarshalTagsexpr(t *testing.T) {
	type testT struct {
		A string
		B string `sexpr:"bTag"`
	}
	ts := []struct {
		s        string
		expected testT
	}{
		{
			expected: testT{
				A: "a",
				B: "bar",
			},
			s: `((A "a") (bTag "bar"))`,
		},
		{
			expected: testT{
				B: "bar",
			},
			s: `( (bTag "bar"))`,
		},
		{
			expected: testT{
				A: "bar",
			},
			s: `((A "bar"))`,
		},
		{
			expected: testT{},
			s:        `()`,
		},
	}

	for _, tc := range ts {
		var got testT
		if err := Unmarshal([]byte(tc.s), &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("unexpected s expr. expected: %q, but got: %q", tc.expected, got)
		}
	}
}

func TestMarshalComplex(t *testing.T) {
	var ts = []struct {
		c        [2]float32
		expected string // expected error from Parse/Check or result from Eval
	}{
		{[2]float32{1, 2}, "#C(1.0 2.0)"},
		{[2]float32{1, 0}, "#C(1.0 0.0)"},
		{[2]float32{0, 1}, "#C(0.0 1.0)"},
		{[2]float32{3, 10}, "#C(3.0 10.0)"},
	}

	for _, tc := range ts {
		c := complex(tc.c[0], tc.c[1])
		b, err := Marshal(c)
		if err != nil {
			t.Fatal(err)
		}

		if got := string(b); got != tc.expected {
			t.Errorf("unexpected bytes. expected: %s, but got: %v", tc.expected, string(got))
		}
	}
}

func TestMarshalBool(t *testing.T) {
	var ts = []struct {
		b        bool
		expected string // expected error from Parse/Check or result from Eval
	}{
		{true, string([]byte{'t'})},
		{false, "nil"},
	}

	for _, tc := range ts {
		b, err := Marshal(tc.b)
		if err != nil {
			t.Fatal(err)
		}

		if got := string(b); got != tc.expected {
			t.Errorf("unexpected bytes. expected: %s, but got: %v", tc.expected, got)
		}
	}
}

func TestMarshalInterface(t *testing.T) {
	var ts = []struct {
		x        interface{}
		expected string // expected error from Parse/Check or result from Eval
	}{
		{
			x:        []int{1, 2, 3},
			expected: `("[]int" (1 2 3))`,
		},
		{
			x:        []string{"a", "b", "c"},
			expected: `("[]string" ("a" "b" "c"))`,
		},
		{
			x:        nil,
			expected: "nil",
		},
	}

	for _, tc := range ts {
		b, err := Marshal(&tc.x)
		if err != nil {
			t.Fatal(err)
		}

		if got := string(b); got != tc.expected {
			t.Errorf("unexpected bytes. expected: %s, but got: %v", tc.expected, got)
		}
	}
}

func TestMarshal(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Oscars          []string
		Sequel          *string
	}

	var ts = []struct {
		m        Movie
		expected string // expected error from Parse/Check or result from Eval
	}{
		{
			m: Movie{
				Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
				Year:     1964,
				Oscars: []string{
					"Best Actor (Nomin.)",
					"Best Adapted Screenplay (Nomin.)",
				},
			},
			expected: "( (Subtitle \"How I Learned to Stop Worrying and Love the Bomb\") (Year 1964) (Oscars (\"Best Actor (Nomin.)\" \"Best Adapted Screenplay (Nomin.)\")))",
		},
		{
			m: Movie{
				Title: "Dr. Strangelove",
				Oscars: []string{
					"Best Actor (Nomin.)",
					"Best Adapted Screenplay (Nomin.)",
				},
			},
			expected: "((Title \"Dr. Strangelove\") (Oscars (\"Best Actor (Nomin.)\" \"Best Adapted Screenplay (Nomin.)\")))",
		},
	}

	for _, tc := range ts {
		b, err := Marshal(tc.m)
		if err != nil {
			t.Fatal(err)
		}

		if got := string(b); got != tc.expected {
			t.Errorf("unexpected bytes. \nexpected: \n%v\nbut got: \n%v", tc.expected, got)
		}
	}
}

func TestUnmarshal(t *testing.T) {
	t.Run("Movie", func(t *testing.T) {
		type Movie struct {
			Title, Subtitle string
			Year            int
			Oscars          []string
			Sequel          *string
		}

		var ts = []struct {
			expected Movie
			m        string // expected error from Parse/Check or result from Eval
		}{
			{
				expected: Movie{
					Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
					Year:     1964,
					Oscars: []string{
						"Best Actor (Nomin.)",
						"Best Adapted Screenplay (Nomin.)",
					},
				},
				m: "( (Subtitle \"How I Learned to Stop Worrying and Love the Bomb\") (Year 1964) (Oscars (\"Best Actor (Nomin.)\" \"Best Adapted Screenplay (Nomin.)\")))",
			},
			{
				expected: Movie{
					Title: "Dr. Strangelove",
					Oscars: []string{
						"Best Actor (Nomin.)",
						"Best Adapted Screenplay (Nomin.)",
					},
				},
				m: "((Title \"Dr. Strangelove\") (Oscars (\"Best Actor (Nomin.)\" \"Best Adapted Screenplay (Nomin.)\")))",
			},
		}

		for _, tc := range ts {
			var got Movie
			if err := Unmarshal([]byte(tc.m), &got); err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("unexpected bytes. \nexpected: \n%v\nbut got: \n%v", tc.expected, got)
			}
		}
	})

	t.Run("Bool", func(t *testing.T) {
		var b bool
		if err := Unmarshal([]byte(`t`), &b); err != nil {
			t.Fatal(err)
		}
		if b != true {
			t.Error("unexpected boolean. expected: true, but got: false")
		}
	})

	t.Run("Complex", func(t *testing.T) {
		var c complex128
		if err := Unmarshal([]byte(`#C(1.0 2.0)`), &c); err != nil {
			t.Fatal(err)
		}
		if c != 1+2i {
			t.Errorf("unexpected complex. expected: 1+2i, but got: %v", c)
		}
	})

	t.Run("Float", func(t *testing.T) {
		var f float64
		if err := Unmarshal([]byte(`2.5`), &f); err != nil {
			t.Fatal(err)
		}
		if f != 2.5 {
			t.Errorf("unexpected float. expected: 2.5, but got: %v", f)
		}
	})

	t.Run("InterfaceSlice", func(t *testing.T) {
		var i interface{}
		if err := Unmarshal([]byte(`("[]int" (1 2 3))`), &i); err != nil {
			t.Fatal(err)
		}
		expected := []int{1, 2, 3}
		if !reflect.DeepEqual(i, expected) {
			t.Errorf("unexpected value. expected: %#v, but got: %#v", expected, i)
		}
	})

	t.Run("InterfaceArray", func(t *testing.T) {
		var i interface{}
		if err := Unmarshal([]byte(`("[6]byte" (5 4 3 2 1 0))`), &i); err != nil {
			t.Fatal(err)
		}
		expected := [6]byte{5, 4, 3, 2, 1, 0}
		if !reflect.DeepEqual(i, expected) {
			t.Errorf("unexpected value. expected: %#v, but got: %#v", expected, i)
		}
	})

	t.Run("InterfaceMap", func(t *testing.T) {
		var i interface{}
		if err := Unmarshal([]byte(`("map[string]int" (("a" 1) ("aa" 2) ("日本語" 3)))`), &i); err != nil {
			t.Fatal(err)
		}
		expected := map[string]int{
			"a":   1,
			"aa":  2,
			"日本語": 3,
		}
		if !reflect.DeepEqual(i, expected) {
			t.Errorf("unexpected value. expected: %#v, but got: %#v", expected, i)
		}
	})

	t.Run("InterfaceMapArrayKey", func(t *testing.T) {
		var i interface{}
		if err := Unmarshal([]byte(`("map[[1]byte]int" (((1) 1) ((2) 2) ((3) 3)))`), &i); err != nil {
			t.Fatal(err)
		}
		expected := map[[1]byte]int{
			[1]byte{1}: 1,
			[1]byte{2}: 2,
			[1]byte{3}: 3,
		}
		if !reflect.DeepEqual(i, expected) {
			t.Errorf("unexpected value. expected: %#v, but got: %#v", expected, i)
		}
	})

	t.Run("InterfaceMapSliceValue", func(t *testing.T) {
		var i interface{}
		if err := Unmarshal([]byte(`("map[int][]int" ((-10000 (1 2 3)) (200000 (4 5 6)) (3000000000 (7 8 9)))`), &i); err != nil {
			t.Fatal(err)
		}
		expected := map[int][]int{
			-10000:     {1, 2, 3},
			200000:     {4, 5, 6},
			3000000000: {7, 8, 9},
		}
		if !reflect.DeepEqual(i, expected) {
			t.Errorf("unexpected value. expected: %#v, but got: %#v", expected, i)
		}
	})

	t.Run("InterfaceMapInterface", func(t *testing.T) {
		var i interface{}
		if err := Unmarshal([]byte(`("map[string]interface{}" (("a" ("[]int" (1 2 3))) ("b" ("string" "hello"))))`), &i); err != nil {
			t.Fatal(err)
		}
		expected := map[string]interface{}{
			"a": []int{1, 2, 3},
			"b": "hello",
		}
		if !reflect.DeepEqual(i, expected) {
			t.Errorf("unexpected value. expected: %#v, but got: %#v", expected, i)
		}
	})

}
