package asparser

import (
	"fmt"
	"log"
)

var LatexConversionTable *ConversionTable = &ConversionTable{
	newLineString: "\n",
	table: map[string]DictionaryToken{
		"block": {
			idName: "block",
			tokenName: "block",
		},
		"times": {
			idName: "times",
			tokenName: "times", 
			converter: func(token Token, ct ConversionTable) string {
				a, err := ct.Convert(*token.content[0])
				if err != nil {
					log.Println("failed to 'a' of times")
				}
				b, err := ct.Convert(*token.content[1])
				if err != nil {
					log.Println("failed to 'b' of times")
				}
				return fmt.Sprintf("(%s \\times %s)", a, b)
			},
		},
		"root": {
			idName: "root",
			tokenName: "root", 
			converter: func(token Token, ct ConversionTable) (string) {
				base := "2"
				if len(token.content) >= 2 && token.content[0].body != "" {
					b, err := ct.Convert(*token.content[0])
					if err != nil {
						log.Println("failed to parse base of root")
					} else {
						base = b
					}
				}
				operand, err := ct.Convert(*token.content[len(token.content)-1])
				if err != nil {
					log.Println("failed to parse operand of root")
				}
				return fmt.Sprintf("\\root[%s](%s)", base, operand)
			},
		},
	},
}
