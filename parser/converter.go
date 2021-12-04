package asparser

type ConversionTable struct {
	table map[string]DictionaryToken
}

type DictionaryToken struct {
	idName string // Should match the key in the map
	tokenName string // The grammar token this corrisponds with
	converter func(token Token) string // Returns the converted string
}

func(ct ConversionTable) Convert (token Token) (result string) {
	if ct.table[token.grammar.idName].converter != nil {
		result += ct.table[token.grammar.idName].converter(token)	
	}
	
	for _, t := range token.content {
		result += ct.Convert(*t)
	}
	return
}