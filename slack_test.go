package main

import (
	"testing"
)

func TestValid(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Name      string
		TestInput string
		Expected  bool
	}{
		{"case1", "a", true},
		{"case2", "", false},
		{"case3", "#a", true},
		{"case4", "@a", true},
	}

	for _, tc := range cases {
		s := SlackChannel{Name: tc.TestInput}

		actual := s.Valid()

		if actual != tc.Expected {
			t.Errorf("[%s] Expected %v but got %v", tc.Name, tc.Expected, actual)
		}
	}
}
