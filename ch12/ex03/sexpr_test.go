package sexpr

import (
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
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

func TestMarchalInterface(t *testing.T) {
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

// func Test(t *testing.T) {
// 	type Movie struct {
// 		Title, Subtitle string
// 		Year            int
// 		Actor           map[string]string
// 		Oscars          []string
// 		Sequel          *string
// 	}
// 	strangelove := Movie{
// 		Title:    "Dr. Strangelove",
// 		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
// 		Year:     1964,
// 		Actor: map[string]string{
// 			"Dr. Strangelove":            "Peter Sellers",
// 			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
// 			"Pres. Merkin Muffley":       "Peter Sellers",
// 			"Gen. Buck Turgidson":        "George C. Scott",
// 			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
// 			`Maj. T.J. "King" Kong`:      "Slim Pickens",
// 		},
// 		Oscars: []string{
// 			"Best Actor (Nomin.)",
// 			"Best Adapted Screenplay (Nomin.)",
// 			"Best Director (Nomin.)",
// 			"Best Picture (Nomin.)",
// 		},
// 	}

// 	// Encode it
// 	data, err := Marshal(strangelove)
// 	if err != nil {
// 		t.Fatalf("Marshal failed: %v", err)
// 	}
// 	t.Logf("Marshal() = %s\n", data)

// 	// Decode it
// 	var movie Movie
// 	if err := Unmarshal(data, &movie); err != nil {
// 		t.Fatalf("Unmarshal failed: %v", err)
// 	}
// 	t.Logf("Unmarshal() = %+v\n", movie)

// 	// Check equality.
// 	if !reflect.DeepEqual(movie, strangelove) {
// 		t.Fatal("not equal")
// 	}

// 	// Pretty-print it:
// 	data, err = MarshalIndent(strangelove)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("MarshalIdent() = %s\n", data)
// }
