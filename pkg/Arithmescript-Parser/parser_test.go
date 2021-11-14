package asparser

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestReplace(t *testing.T) {
	s := "start body end start body part1, part2 end"
	grammar := &Grammar{}

	grammar.rules = append(grammar.rules, &GrammarToken{"start", ",", "end", "hi", "#", "bye"})
	// grammar.rules = append(grammar.rules, &GrammarToken{"\root", "of", ")", "hi", ",", "bye"})

	want := "hi body bye hi body part1# part2 bye"
	msg, err := ParseString(s, grammar)
	if want != msg || err != nil {
		t.Fatalf(`parseString(s, grammar) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
