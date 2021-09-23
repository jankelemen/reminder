package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		name, input string
		expected    []string
	}{
		{name: "Unpack 1 digit", input: "21-20-1{0,1,2}", expected: []string{"21-20-10", "21-20-11", "21-20-12"}},
		{name: "Unpack 2 digits", input: "21-{20,21}-10", expected: []string{"21-20-10", "21-21-10"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Unpack(test.input)
			assertStringArrays(t, got, test.expected)
		})
	}
}

func TestUnpackAll(t *testing.T) {
	input := "2x-{1}{0,2}-{20,21}"
	got := UnpackAll(input)
	expected := []string{"2x-10-20", "2x-10-21", "2x-12-20", "2x-12-21"}
	assertStringArrays(t, got, expected)
}

func TestIsTheSameTime(t *testing.T) {
	tests := []struct {
		name, time, timeReadFromFile string
		expected                     bool
	}{
		{name: "Without wildcards", time: "21-10-20", timeReadFromFile: "21-10-20", expected: true},
		{name: "With month as wildcard", time: "21-10-10", timeReadFromFile: "21-xx-10", expected: true},
		{name: "With day as wildcard", time: "21-10-10", timeReadFromFile: "21-10-xx", expected: true},
		{name: "Wrong input with wildcard", time: "21-10-21", timeReadFromFile: "xx-10-x0", expected: false},
		{name: "Wrong input without wildcard", time: "21-10-10", timeReadFromFile: "21-10-20", expected: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsTheSameTime(test.time, test.timeReadFromFile)
			if test.expected != got {
				t.Errorf("time: %s, timeReadFromFile: %s, got: %t, want: %t", test.time, test.timeReadFromFile, got, test.expected)
			}
		})
	}
}

func TestCurrentDate(t *testing.T) {
	got := CurrentDate()
	if strings.Count(got, "-") != 2 || len(got) != 8 {
		t.Errorf("date in wrong format: %s", got)
	}
}

func assertStringArrays(t testing.TB, got, want []string) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %#v, want: %#v", got, want)
	}
}
