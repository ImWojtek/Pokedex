package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "basic test",
			input:    "      hello     world    ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "mixed case test",
			input:    "pikachu CHARMANDER Squirtle ",
			expected: []string{"pikachu", "charmander", "squirtle"},
		},
		{
			name:     "comma separator",
			input:    " bulbasaur, ivysaur, venusaur ",
			expected: []string{"bulbasaur", "ivysaur", "venusaur"},
		},
		{
			name:     "comma no whitespace",
			input:    "eevee,jolteon,flareon",
			expected: []string{"eevee", "jolteon", "flareon"},
		},
	}

	for _, c := range cases {
		got := cleanInput(c.input)
		if !reflect.DeepEqual(c.expected, got) {
			t.Fatalf("%s: expected %v, got %v", c.name, c.expected, got)
		}
	}
}
