package asparser

import (
	"fmt"
	"testing"
)

func TestConverter(t *testing.T) {
	t.Run("nested tokens with ASGRAMMAR", func(t *testing.T) {
		toks, err := ParseMultiline("(a times b) times (c times d)", ASGRAMMAR)
		if err != nil { t.Fail() }
		fmt.Println(toks)
		lat, err := LatexConversionTable.Convert(toks[0])
		println(lat)
		if err != nil { t.Fail() }
		if lat != `(a \times b) \times (c \times d)` { t.Fail() }
	})

	t.Run("nested times in root with ASGRAMMAR", func(t *testing.T) {
		toks, err := ParseMultiline("8th root 256 times 7;", ASGRAMMAR)
		fmt.Println(toks)
		if err != nil { t.Fail() }
		lat, err := LatexConversionTable.ConvertMultiline(toks)
		println(lat)
		if err != nil { t.Fail() }
		if lat != `\root[8]{256 \times 7}` { t.Fail() }
	})
	t.Run("nested root in root with ASGRAMMAR", func(t *testing.T) {
		toks, err := ParseMultiline("8th root root 64;", ASGRAMMAR)
		fmt.Println(toks)
		if err != nil { t.Fail() }
		lat, err := LatexConversionTable.ConvertMultiline(toks)
		println(lat)
		if err != nil { t.Fail() }
		if lat != `\root[8](\root[2]{64})` { t.Fail() }
	})
	t.Run("sqrt", func(t *testing.T) {
		toks, err := ParseMultiline("root of 16; ", ASGRAMMAR)
		fmt.Println(toks)
		if err != nil { t.Fail() }
		lat, err := LatexConversionTable.ConvertMultiline(toks)
		println(lat)
		if err != nil { t.Fail() }
		if lat != `\root[2]{16}` { t.Fail() }
	})
	t.Run("sqrt", func(t *testing.T) {
		toks, err := ParseMultiline("root 2 of 16; ", ASGRAMMAR)
		fmt.Println(toks)
		if err != nil { t.Fail() }
		lat, err := LatexConversionTable.ConvertMultiline(toks)
		println(lat)
		if err != nil { t.Fail() }
		if lat != `\root[2]{16}` { t.Fail() }
	})
	
	// TODO: Write MORE tests

}