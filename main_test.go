package main

import (
	"testing"
)

func TestParseSlackChannels(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Name        string
		ChannelList string
		Expected    []string
	}{
		{"case1", "a", []string{"#a"}},
		{"case2", "@a", []string{"@a"}},
		{"case3", "#a", []string{"#a"}},
		{"case4", "a,b", []string{"#a", "#b"}},
		{"case5", "#a,b", []string{"#a", "#b"}},
		{"case6", "#a,b", []string{"#a", "#b"}},
		{"case7", "a,#b", []string{"#a", "#b"}},
		{"case8", "#a,#b", []string{"#a", "#b"}},
		{"case9", "", []string{}},
		{"case10", "           ", []string{}},
		{"case11", " ,   ,   , ,, ", []string{}},
		{"case12", " , c ,   ,a,, ", []string{"#c", "#a"}},
		{"case13", "a,#b,@c", []string{"#a", "#b", "@c"}},
	}

	for _, tc := range cases {
		actual := parseSlackChannels(tc.ChannelList)

		if !stringSliceEq(actual, tc.Expected) {
			t.Errorf("[%s] Expected %v but got %v", tc.Name, tc.Expected, actual)
		}
	}
}

func stringSliceEq(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
