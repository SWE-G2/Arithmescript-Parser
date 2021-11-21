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

	t.Run("anyword surround by stuff", func(t *testing.T) { 	
		if PickGrammarForLine("loremanywordipsum", grammar) != grammar.rules[0] {
			t.Fail()
		} 
	})
	t.Run("anyword surround by nothing", func(t *testing.T) { 	
		if PickGrammarForLine("anyword", grammar) != grammar.rules[0] {
			t.Fail()
		} 
	})
	t.Run("oneword surround by nothing", func(t *testing.T) { 	
		if PickGrammarForLine("oneword", grammar) != grammar.rules[1] {
			t.Fail()
		} 
	})
	t.Run("oneword surround by stuff", func(t *testing.T) { 	
		if PickGrammarForLine("loremonewordipsum", grammar) != nil {
			t.Fail()
		} 
	})
	t.Run("undefined token", func(t *testing.T) { 	
		if PickGrammarForLine("owo", grammar) != nil {
			t.Fail()
		} 
	})
	t.Run("times surround by stuff", func(t *testing.T) { 	
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