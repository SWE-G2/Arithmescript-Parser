package asparser

import (
	"testing"
)

func TestConverter(t *testing.T) {
	t.Run("nested tokens with ASGRAMMAR", func(t *testing.T) {
		toks, err := ParseMultiline("(a times b) times (c times d)", ASGRAMMAR)
		if err != nil { t.Fail() }
		lat, err := LatexConversionTable.Convert(toks[0])
		println(lat)
		if err != nil { t.Fail() }
		if lat != `(a \times b) \times (c \times d)` { t.Fail() }
	})

	t.Run("nested tokens with ASGRAMMAR #2", func(t *testing.T) {
		toks, err := ParseMultiline("8th root 256 times 7; root of 16; ", ASGRAMMAR)
		// fmt.Println(toks)
		if err != nil { t.Fail() }
		lat, err := LatexConversionTable.ConvertMultiline(toks)
		println(lat)
		if err != nil { t.Fail() }
		if lat != `\root[8]((256 \times 7))\n\root[2](16)` { t.Fail() }
	})
	
	// TODO: Write MORE tests

}