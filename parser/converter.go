package asparser

type ConversionTable struct {
	table map[string]DictionaryToken
	newLineString string // What should end each line?
}

type DictionaryToken struct {
	idName string // Should match the key in the map
	tokenName string // The grammar token this corrisponds with
	converter func(token Token, ct ConversionTable) string // Returns the converted string
}

func(ct ConversionTable) Convert (token Token) (result string, err error) {
	result = token.body
	if token.grammar == nil {
		return
	}

	if _, exists := ct.table[token.grammar.idName]; !exists { // Check if map value exists
		return 
	}
	
	if ct.table[token.grammar.idName].converter != nil { 
		result = ct.table[token.grammar.idName].converter(token, ct)
		return
	}
	result = ""
	
	for _, t := range token.content { 
		s, err := ct.Convert(*t)
		if err != nil {
			return result, err
		}
		result += s
	}
	return
}

func(ct ConversionTable) ConvertMultiline (tokens []Token) (result string, err error)  {
	for _, t := range tokens {
		s, e := ct.Convert(t)
		if e != nil {
			err = e
			return
		}
		result += s + ct.newLineString
	}
	return
}