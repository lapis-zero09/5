package main

import (
	"encoding/xml"
	"testing"
)

func TestGetSelectors(t *testing.T) {
	ts := []struct {
		selectors SelectorList
		expected  string
	}{
		{
			selectors: SelectorList{
				&TypeSelector{"div"},
			},
			expected: "div",
		},
		{
			selectors: SelectorList{
				&AttrSelector{"class", "className"},
			},
			expected: "className",
		},
		{
			selectors: SelectorList{
				&AttrSelector{"id", "idName"},
			},
			expected: "idName",
		},
		{
			selectors: SelectorList{
				&TypeSelector{"div"},
				&TypeSelector{"div"},
				&TypeSelector{"h2"},
			},
			expected: "div div h2",
		},
		{
			selectors: SelectorList{
				&TypeSelector{"div"},
				&AttrSelector{"id", "idName"},
				&AttrSelector{"class", "className"},
				&TypeSelector{"p"},
			},
			expected: "div idName className p",
		},
	}

	for _, tc := range ts {
		if got := tc.selectors.GetSelectors(); got != tc.expected {
			t.Errorf("unexpected result. expected: %v, but got: %v", tc.expected, got)
		}
	}
}

func TestArgParse(t *testing.T) {
	ts := []struct {
		args     []string
		expected string
	}{
		{
			args: []string{"div"},
			expected: SelectorList{
				&TypeSelector{"div"},
			}.GetSelectors(),
		},
		{
			args: []string{"#className"},
			expected: SelectorList{
				&AttrSelector{"class", "className"},
			}.GetSelectors(),
		},
		{
			args: []string{".idName"},
			expected: SelectorList{
				&AttrSelector{"id", "idName"},
			}.GetSelectors(),
		},
		{
			args: []string{"div", "div", "h2"},
			expected: SelectorList{
				&TypeSelector{"div"},
				&TypeSelector{"div"},
				&TypeSelector{"h2"},
			}.GetSelectors(),
		},
		{
			args: []string{"div", "idName", "className", "p"},
			expected: SelectorList{
				&TypeSelector{"div"},
				&AttrSelector{"id", "idName"},
				&AttrSelector{"class", "className"},
				&TypeSelector{"p"},
			}.GetSelectors(),
		},
	}

	for _, tc := range ts {
		got, err := argParse(tc.args)
		if err != nil {
			t.Errorf("Error has occured: %v", err)
		}
		if got.GetSelectors() != tc.expected {
			t.Errorf("unexpected result. expected: %v, but got: %v", tc.expected, got)
		}
	}
}

func TestContainsAll(t *testing.T) {
	ts := []struct {
		elems     []xml.StartElement
		selectors SelectorList
		expected  bool
	}{
		{
			elems: []xml.StartElement{
				{Attr: []xml.Attr{{xml.Name{Local: "class"}, "foo"}}},
			},
			selectors: []Selector{
				&AttrSelector{"class", "foo"},
			},
			expected: true,
		},
		{
			elems: []xml.StartElement{
				{},
				{},
				{Attr: []xml.Attr{{xml.Name{Local: "id"}, "bar"}}},
				{},
				{},
			},
			selectors: []Selector{
				&AttrSelector{"id", "bar"},
			},
			expected: true,
		},
	}

	for _, tc := range ts {
		if got := containsAll(tc.elems, tc.selectors); got != tc.expected {
			t.Errorf("unexpected result. expected: %v, but got: %v", tc.expected, got)
		}
	}
}
