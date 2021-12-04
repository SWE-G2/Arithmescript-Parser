package asparser

import (
	"regexp"
	"strings"
)

var ASGRAMMAR *Grammar = &Grammar{
	rules: map[string]*GrammarToken{
		"block": {
			idName: "block",
			keywords: *regexp.MustCompile(`\((?:[^)(]+)*\)`), // Should NOT be recusive! 
			decomposer: func(self *GrammarToken, asMarkup string) (result []string) {
				result = self.keywords.FindAllString(asMarkup, -1)
				for i, s := range result {
					result[i] = strings.Trim(s, "()")
				}
				return
			},
		},
		"root": {
			idName: "root",
			keywords: *regexp.MustCompile("root"), 
			decomposer: func(self *GrammarToken, asMarkup string) (result []string) {
				splitAtRoot := strings.SplitN(asMarkup, "root", 2)
				splitAtOf := strings.SplitN(splitAtRoot[len(splitAtRoot)-1], "of", 2)
				
				if len(splitAtRoot) == 2 {
					splitAtRoot[0] = strings.TrimSpace(splitAtRoot[0])
					splitAtRoot[0] = strings.TrimSuffix(splitAtRoot[0], "th")
					splitAtRoot[0] = strings.TrimSuffix(splitAtRoot[0], "st")
					splitAtRoot[0] = strings.TrimSuffix(splitAtRoot[0], "nd")
					result = append(result, splitAtRoot[0])
				}
				last := len(splitAtOf)-1
				splitAtOf[last] = strings.TrimSpace(splitAtOf[last])
				splitAtOf[last] = strings.TrimPrefix(splitAtOf[last], "of")
				result = append(result, splitAtOf[last])
				return
		}},
		"times": {
			idName: "times",
			keywords: *regexp.MustCompile("times"), 
			decomposer: func(self *GrammarToken, asMarkup string) []string {
			return strings.Split(asMarkup, "times")
		}},

	}}
