package sentiment

import (
	"testing"
)

func TestTokenizeSimple(t *testing.T) {
	tokens := tokenize("A simple sentence")

	if len(tokens) != 3 {
		t.Errorf("We should have 3 tokens but %d given. %v", len(tokens), tokens)
	}

	expected := []string{"A", "simple", "sentence"}
	for i, v := range tokens {
		if v != expected[i] {
			t.Errorf("Extracted tokens are not valid: %s != %s", tokens[i], expected[i])
		}
	}
}

func TestTokenizeStrange(t *testing.T) {
	tokens := tokenize("This is a (strange) sentence!")

	if len(tokens) != 5 {
		t.Errorf("We should have 5 tokens but %d given. %v", len(tokens), tokens)
	}

	expected := []string{"This", "is", "a", "strange", "sentence"}
	for i, v := range tokens {
		if v != expected[i] {
			t.Errorf("Extracted tokens are not valid: %s != %s", tokens[i], expected[i])
		}
	}
}
