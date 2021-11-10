package asparser

import (
	"strings"
)

type GrammarToken struct {
	asStart        string
	asDelimiter    string
	asEnd          string
	latexStart     string
	latexDelimiter string
	latexEnd       string
}

type Grammar struct {
	rules []*GrammarToken
}

type Token struct {
	body    string
	content []*Token
	grammar *GrammarToken
}

// Maybe we should drop tokenization? I don't think we actually need it unless we 
// make more than one target language.

// func parseStringByWords(asMarkup string, grammar *Grammar) (parsedTokens []Token, error) {
// 	words := strings.Split(asMarkup, " ")
// 	currToken := Token
// 	hasFoundStart := false
// 	hasFoundEnd := false
// 	hasFoundBody := false
// 	hasFoundDelimiter := false
// 	for wordIndex, word := words {

// 		for _, rule := range grammar.rules {
// 			if strings.Contains(word, rule.asStart) {
// 				Token.grammar = &rule
// 				parsedTokens = append(parsedTokens, t)
// 				continue

// 			}
// 			if strings.Contains(word, rule.asDelimiter) {

// 				continue
// 			}
// 		}
// 		if hasFoundStart && hasFoundEnd && !hasFoundDelimiter {
// 			currToken = words[wordIndex - 1]
// 		}
// 	}
// 	return parsedTokens, nil
// }

func parseString(asMarkup string, grammar *Grammar) (s string, err error) {
	words := strings.Split(asMarkup, " ")
	for _, word := range words {
		for _, rule := range grammar.rules {
			if strings.Contains(word, rule.asStart) {
				s = s + rule.latexStart + " "
				continue
			}

			if strings.Contains(word, rule.asDelimiter) {
				s = s + strings.TrimRight(word, rule.asDelimiter) + rule.latexDelimiter + " "
				continue
			}

			if strings.Contains(word, rule.asEnd) {
				s = s + rule.latexEnd + " "
				continue
			}
			s += word + " "
		}
	}
	s = strings.Trim(s, " ")
	return
}
