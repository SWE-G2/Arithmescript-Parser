package asparser

var LatexConversionTable *ConversionTable = &ConversionTable{
	table: map[string]DictionaryToken{
		"block": {
			idName: "block",
			tokenName: "block",
		},
		"times": {
			idName: "times",
			tokenName: "times", 
			converter: func(token Token) string {
				return `\times`
			},
		},
		"root": {
			idName: "root",
			tokenName: "root", 
			converter: func(token Token) (string) {
				return `\root`
			},
		},
	},
}
