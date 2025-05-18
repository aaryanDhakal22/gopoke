package utils

import (
	"testing"
)

type Case struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {

	items := []Case{
		{
			input:    " hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmandar BulbaSaur PIKACHU",
			expected: []string{"charmandar", "bulbasaur", "pikachu"},
		},
	}

	for _, item := range items {
		t.Run(item.input, func(t *testing.T) {
			actual := CleanInput(item.input)
			if len(actual) != len(item.expected) {
				t.Fatalf("Expected length of array was %v but got %v", len(item.expected), len(actual))
			}
			for j := range actual {
				if actual[j] != item.expected[j] {
					t.Fatalf("Expected : %v Received : %v", item.expected[j], actual[j])
				}
			}
		})
	}
}
