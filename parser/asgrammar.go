package asparser

import (
	"regexp"
	"strings"
)
var ASGRAMMAR *Grammar = &Grammar{
	rules: map[string]*GrammarToken{
		"times": {keywords: *regexp.MustCompile("times"), decomposer: func(asMarkup string) []string {
			return strings.Split(asMarkup, "times")
		}},
		"root": {keywords: *regexp.MustCompile("root"), decomposer: func(asMarkup string) (result []string) {
			splitAtRoot := strings.SplitN(asMarkup, "root", 2)
			splitAtOf := strings.SplitN(splitAtRoot[len(splitAtRoot)-1], "of", 2)
			if len(splitAtRoot) == 2 {
				result = append(result, splitAtRoot[0])
			}
			result = append(result, splitAtOf[len(splitAtOf)-1])
			return
		}},
}}