package main

import (
	"fmt"
	"syscall/js"

	asp "github.com/SWE-G2/Arithmescript-Parser/parser"
)


func sayHello(self js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		js.Global().Get("console").Call("error", "Invalid number of arguments")
		return nil
	}
    fmt.Println("Hello ", args[0])
	return nil
}

// // Parse a string with ASGRAMMER
// func ParseMultilineASG(self js.Value, args []js.Value) interface{} {
// 	asp.ParseMultiline(args[0].String(), asp.ASGRAMMAR)
// 	return nil
// }

// Convert AS to Latex
func ConvertASToLatex(self js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		js.Global().Get("console").Call("error", "Invalid number of arguments")
		return nil
	}
	tokens, err := asp.ParseMultiline(args[0].String(), asp.ASGRAMMAR)
	if err != nil {
		return err
	}
	var latexString string
	s, err := asp.LatexConversionTable.ConvertMultiline(tokens)
	if err != nil {
		return err
	}
	latexString = s

	return latexString
}

func main() {  
    fmt.Println("Go Web Assembly")
    // js.Global().Set("ParseMultilineASG", js.FuncOf(ParseMultilineASG))
    js.Global().Set("sayHello", js.FuncOf(sayHello))
    js.Global().Set("ConvertASToLatex", js.FuncOf(ConvertASToLatex))
	<-make(chan bool)	// Stops program from exiting, 
						// makes go wait for this channel to exit 
						// (it never will, muhuhahahaha!)
}
// Use this to compile for WASM
// GOOS=js GOARCH=wasm go build -o asparser.wasm