package asparser

type ConversionTable struct {
	table map[string]string
}

func(ct ConversionTable) Convert (token Token, args ...string) (out string) {
	// target := ct.table[token.grammar]
	// out = fmt.Sprintf(, args)
	return
}