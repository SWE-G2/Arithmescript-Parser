package asparser

import (
	"regexp"
	"strings"
)

type GrammarToken struct {
	keywords   regexp.Regexp // The regex use to match a line to this GT. Use regexp.MustCompile to define this.
	definition string // A string that gets-- Wait, I don't remember why I added this, conversions will go in a new module.
	decomposer func(asMarkup string) ([]string) // Returns a list of strings to be converted to tokens.
}

type Grammar struct {
	rules []*GrammarToken
}

type Token struct {
	body    string
	content []*Token
	grammar *GrammarToken
}


func PickGrammarForLine(asMarkup string, grammar *Grammar) (t *GrammarToken) {
	for _, grammarToken := range grammar.rules {
		regexResult := grammarToken.keywords.Find([]byte(asMarkup))
		if len(regexResult) != 0 {
			return grammarToken;
		}
	}
	return nil
}

// Parses a single line of AS markup
func ParseExpression(asMarkup string, grammar *Grammar) (parsed Token, err error)  {
		asMarkup = strings.TrimSpace(asMarkup)
		asMarkup = strings.Trim(asMarkup, "{}()[]")
		asMarkup = strings.ToLower(asMarkup)
		
		grammarToken := PickGrammarForLine(asMarkup, grammar)
		
		parsed.body = asMarkup

		if grammarToken == nil {
			return
		}
		parsed.grammar = grammarToken
		
		if grammarToken.decomposer == nil {
			return
		}
		for _, s := range grammarToken.decomposer(asMarkup) {
			child, childError := ParseExpression(s, grammar)
			
			if childError != nil {
				err = childError
				return
			}
			
			parsed.content = append(parsed.content, &child)
		}
		return 
}

// Parses multiple lines of AS markup
func ParseMultiline(asMarkup string, grammar *Grammar) (parsed []Token, err error) {
	lines := strings.Split(asMarkup, ";")
	for _, line := range lines {
		go ParseExpression(line, grammar)
	}

	return
}
