package asparser

import (
	"regexp"
	"strings"
	"sync"
)

type GrammarToken struct {
	idName string // Should match the name in the map
	keywords   regexp.Regexp                  // The regex use to match a line to this GT. Use regexp.MustCompile to define this.
	decomposer func(self *GrammarToken, asMarkup string) []string // Returns a list of strings to be converted to tokens.
}

type Grammar struct {
	rules map[string]*GrammarToken
}

type Token struct {
	body    string
	content []*Token
	grammar *GrammarToken
}

func (token Token) String() string {
	return sPrintTree(token, 0)
}

func PickGrammarForLine(asMarkup string, grammar *Grammar) (t *GrammarToken) {
	for _, grammarToken := range grammar.rules {
		regexResult := grammarToken.keywords.Find([]byte(asMarkup))
		if len(regexResult) != 0 {
			return grammarToken
		}
	}
	return nil
}

// Parses a single line of AS markup
func ParseExpression(asMarkup string, grammar *Grammar) (parsed Token, err error) {
	asMarkup = strings.TrimSpace(asMarkup)
	// asMarkup = strings.Trim(asMarkup, "{}()[]") // Replaced with order of operations
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
	for _, s := range grammarToken.decomposer(grammarToken, asMarkup) {
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
	lineSpliter := regexp.MustCompile(";|\n")

	lines := lineSpliter.Split(asMarkup, -1)

	wg := sync.WaitGroup{}
	wg.Add(len(lines))

	parsed = make([]Token, len(lines))
	for lineNum, line := range lines {
		go func(lineNum int, line string) {
			defer wg.Done()
			tok, err := ParseExpression(line, grammar)
			if err != nil {
				return
			}
			parsed[lineNum] = tok

		}(lineNum, line)
	}
	wg.Wait()
	return
}

func sPrintTree(token Token, depth int) (result string){
	for i := 0; i < depth; i++ {
		result += "\t"
	}
	result += token.body  + "\n"
	for _, t := range token.content {
		result += sPrintTree(*t, depth + 1)
	}
	return
}