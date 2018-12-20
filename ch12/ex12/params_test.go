package params

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestUnpackValid(t *testing.T) {
	type testType struct {
		A          string `http:"a"`
		CreditCard string `http:"cc" valid:"creditcard"`
	}
	ts := []struct {
		i        map[string]string
		expected testType
	}{
		{
			i: map[string]string{
				"a": "bbb",
			},
			expected: testType{A: "bbb"},
		},
		{
			i: map[string]string{
				"a":  "bbb",
				"cc": "5248012345678901",
			},
			expected: testType{
				A:          "bbb",
				CreditCard: "5248012345678901", // mastercard
			},
		},
		{
			i: map[string]string{
				"a":  "bbb",
				"cc": "3540123456789012",
			},
			expected: testType{
				A:          "bbb",
				CreditCard: "3540123456789012", // jcb
			},
		},
		{
			i: map[string]string{
				"a":  "bbb",
				"cc": "377781234512345",
			},
			expected: testType{
				A:          "bbb",
				CreditCard: "377781234512345", // amex
			},
		},
	}

	for _, tc := range ts {
		var got testType
		req := &http.Request{}
		req.Form = url.Values{}
		for k, v := range tc.i {
			req.Form.Add(k, v)
		}

		if err := Unpack(req, &got); err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("expected: %v, but got: %v", tc.expected, got)
		}
	}
}

func TestUnpackInvalid(t *testing.T) {
	type testType struct {
		A          string `http:"a"`
		CreditCard string `http:"cc" valid:"creditcard"`
	}
	ts := []struct {
		i map[string]string
	}{
		{
			i: map[string]string{
				"cc": "",
			},
		},
		{
			i: map[string]string{
				"cc": "111",
			},
		},
		{
			i: map[string]string{
				"cc": "1111111111111",
			},
		},
		{
			i: map[string]string{
				"cc": "11111111111111",
			},
		},
		{
			i: map[string]string{
				"cc": "111111111111111",
			},
		},
		{
			i: map[string]string{
				"cc": "3777812345123451",
			},
		},
		{
			i: map[string]string{
				"cc": "4540123456789012",
			},
		},
	}

	for _, tc := range ts {
		var got testType
		req := &http.Request{}
		req.Form = url.Values{}
		for k, v := range tc.i {
			req.Form.Add(k, v)
		}

		if err := Unpack(req, &got); err == nil {
			t.Fatal("must be error")
		}
	}
}

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
