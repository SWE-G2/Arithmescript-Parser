package asparser

import (
	"regexp"
	"testing"
)


func TestPickGrammarForLine(t *testing.T) {
	grammar := &Grammar{
		rules: []*GrammarToken{
			{keywords: *regexp.MustCompile("anyword"), definition: "anyword"},
			{keywords: *regexp.MustCompile("^oneword$"), definition: "oneword"},
		},
	}

	t.Run("anyword surrounded by stuff", func(t *testing.T) { 	
		if PickGrammarForLine("loremanywordipsum", grammar) != grammar.rules[0] {
			t.Fail()
		} 
	})
	t.Run("anyword surrounded by nothing", func(t *testing.T) { 	
		if PickGrammarForLine("anyword", grammar) != grammar.rules[0] {
			t.Fail()
		} 
	})
	t.Run("oneword surrounded by nothing", func(t *testing.T) { 	
		if PickGrammarForLine("oneword", grammar) != grammar.rules[1] {
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
		if PickGrammarForLine("a times b", ASGRAMMAR) != ASGRAMMAR.rules[0] {
			t.Fail()
		} 
	})
}

func TestParseExpression(t *testing.T) {
	grammar := &Grammar{
		rules: []*GrammarToken{
			{keywords: *regexp.MustCompile("anyword"), definition: "anyword"},
			{keywords: *regexp.MustCompile("^oneword$"), definition: "oneword"},
		},
	}
	
	t.Run("times with ASGRAMMAR", func(t *testing.T) {
		tok, err := ParseExpression("a times b", ASGRAMMAR)
		if err != nil || (len(tok.content) == 2 && tok.grammar == ASGRAMMAR.rules[1]) {
			t.Fail()
		} 
	})
	t.Run("anyword surrounded by stuff", func(t *testing.T) {
		tok, err := ParseExpression("anyword", grammar)
		if err != nil || tok.grammar != grammar.rules[0] {
			t.Fail()
		} 
	})
}

func TestMultilineParse(t *testing.T) {
		
	t.Run("times with ASGRAMMAR and semicolons", func(t *testing.T) {
		toks, err := ParseMultiline("a times b; c times d", ASGRAMMAR)
		if err != nil || !(len(toks) == 2 && toks[0].grammar == ASGRAMMAR.rules[1] && toks[1].grammar == ASGRAMMAR.rules[1]) {
			t.Fail()
		} 
	})

	t.Run("times with ASGRAMMAR and newlines", func(t *testing.T) {
		toks, err := ParseMultiline("a times b\n c times d", ASGRAMMAR)
		if err != nil || !(len(toks) == 2 && toks[0].grammar == ASGRAMMAR.rules[1] && toks[1].grammar == ASGRAMMAR.rules[1]) {
			t.Fail()
		} 
	})
}