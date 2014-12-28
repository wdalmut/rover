package sentiment

import (
	"testing"
)

func TestExtractFromFiles(t *testing.T) {
	slices := extractFromFile("files/simple.txt")

	if len(slices) != 10 {
		t.Errorf("We should have 10 tokens but %d given. %v", len(slices), slices)
	}

	expected := []string{
		"This", "is", "a", "strange", "sentence",
		"and", "this", "is", "another", "sentence",
	}
	for i, v := range slices {
		if v != expected[i] {
			t.Errorf("Extracted tokens are not valid: %s != %s", slices[i], expected[i])
		}
	}
}
