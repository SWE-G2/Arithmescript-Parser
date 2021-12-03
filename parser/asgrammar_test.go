package asparser

import "testing"

func TestASGrammarDeclaration(t *testing.T) {

	t.Run("ensure idnames match map key", func(t *testing.T) {
		for i, g := range ASGRAMMAR.rules {
			if ASGRAMMAR.rules[g.idName] != g || g.idName != i {
				println("Hey! Rule with idName:", g.idName, " and key:", i, " does not match map key!")
				t.Fail()
			}
		}
	})
}