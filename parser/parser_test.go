package asparser

import (
	"fmt"
	"regexp"
	"testing"
)

func TestPickGrammarForLine(t *testing.T) {
	grammar := &Grammar{
		rules: map[string]*GrammarToken{
			"anyword": {keywords: *regexp.MustCompile("anyword")},
			"oneword": {keywords: *regexp.MustCompile("^oneword$")},
		},
	}

	t.Run("anyword surrounded by stuff", func(t *testing.T) {
		if PickGrammarForLine("loremanywordipsum", grammar) != grammar.rules["anyword"] {
			t.Fail()
		}
	})
	t.Run("anyword surrounded by nothing", func(t *testing.T) {
		if PickGrammarForLine("anyword", grammar) != grammar.rules["anyword"] {
			t.Fail()
		}
	})
	t.Run("oneword surrounded by nothing", func(t *testing.T) {
		if PickGrammarForLine("oneword", grammar) != grammar.rules["oneword"] {
			t.Fail()
		}
	})
	t.Run("oneword surrounded by stuff", func(t *testing.T) {
		if PickGrammarForLine("loremonewordipsum", grammar) != nil {
			t.Fail()
		}
	})
	t.Run("undefined token", func(t *testing.T) {
		if PickGrammarForLine("owo", grammar) != nil {
			t.Fail()
		}
	})
	t.Run("times surrounded by stuff", func(t *testing.T) {
		if PickGrammarForLine("a times b", ASGRAMMAR) != ASGRAMMAR.rules["times"] {
			t.Fail()
		}
	})
}

func TestParseExpression(t *testing.T) {
	grammar := &Grammar{
		rules: map[string]*GrammarToken{
			"anyword": {keywords: *regexp.MustCompile("anyword")},
			"oneword": {keywords: *regexp.MustCompile("^oneword$")},
		},
	}

	t.Run("times with ASGRAMMAR", func(t *testing.T) {
		tok, err := ParseExpression("a times b", ASGRAMMAR)
		if err != nil || !(len(tok.content) == 2 && tok.grammar == ASGRAMMAR.rules["times"]) {
			t.Fail()
		}
	})
	t.Run("anyword surrounded by stuff", func(t *testing.T) {
		tok, err := ParseExpression("anyword", grammar)
		if err != nil || tok.grammar != grammar.rules["anyword"] {
			t.Fail()
		}
	})
	t.Run("nested tokens with ASGRAMMAR", func(t *testing.T) {
		tok, err := ParseExpression("(a times b) times (c times d)", ASGRAMMAR)
		if err != nil { t.Fail() }
		if tok.grammar != ASGRAMMAR.rules["block"] { t.Fail() }
		if tok.content[0].body != "a times b" { t.Fail() }
		if tok.content[1].body != "c times d"  { t.Fail() }
	})
}

func TestMultilineParse(t *testing.T) {

	t.Run("times with ASGRAMMAR and semicolons", func(t *testing.T) {
		toks, err := ParseMultiline("a times b; c times d", ASGRAMMAR)
		if err != nil || !(len(toks) == 2 && toks[0].grammar == ASGRAMMAR.rules["times"] && toks[1].grammar == ASGRAMMAR.rules["times"]) {
			t.Fail()
		}
	})

	t.Run("times with ASGRAMMAR and newlines", func(t *testing.T) {
		toks, err := ParseMultiline("a times b\n c times d", ASGRAMMAR)
		if err != nil || !(len(toks) == 2 && toks[0].grammar == ASGRAMMAR.rules["times"] && toks[1].grammar == ASGRAMMAR.rules["times"]) {
			t.Fail()
		}
	})
	var longString string
	for i := 0; i < 3000; i++ {
		longString += fmt.Sprintf("%v times 1/%v\n", i, i)
	}
	t.Run("times with ASGRAMMAR and newlines with lots of lines", func(t *testing.T) {
		toks, err := ParseMultiline(longString, ASGRAMMAR)
		if err != nil || !(len(toks) == 3001 && toks[1372].grammar == ASGRAMMAR.rules["times"] && toks[290].grammar == ASGRAMMAR.rules["times"]) {
			t.Fail()
		}
	})
}

func printTree(token Token) {
	println(token.grammar)
}
