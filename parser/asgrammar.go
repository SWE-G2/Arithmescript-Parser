package asparser

import (
	"regexp"
	"strings"
)

var asgBlock = &GrammarToken{
			idName: "block",
			keywords: *regexp.MustCompile(`\((?:[^)(]+)*\)`), // Should NOT be recusive! 
			decomposer: func(self *GrammarToken, asMarkup string) (result []string) {
				result = self.keywords.FindAllString(asMarkup, -1)
				for i, s := range result {
					result[i] = strings.Trim(s, "()")
				}
				return
			},
		}

var ASGRAMMAR *Grammar = &Grammar{
	blockDefinition: asgBlock,
	rules: map[string]*GrammarToken{
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
				return strings.Split(skipBlocks(asMarkup), "times")
			},
		},
		"block": asgBlock,
	}}

func splitSkippingBlocks(asMarkup string, splitBy string) (string) {
	
	return asMarkup
}

func splitSkippingBlocksN(asMarkup string) (string) {
	return asMarkup
}
	
	
func skipBlocks(asMarkup string) (string) {
	blocks := asgBlock.decomposer(asgBlock, asMarkup)
	for _, b := range blocks {
		asMarkup = strings.ReplaceAll(asMarkup, b, "")
	}
	return asMarkup
}