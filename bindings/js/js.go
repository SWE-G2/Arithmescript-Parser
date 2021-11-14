package main

import (
	"fmt"
	"syscall/js"
)


func sayHello(self js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		js.Global().Get("console").Call("error", "Invalid number of arguments")
		return nil
	}
    fmt.Println("Hello ", args[0])
	return nil
}

func main() {  
    fmt.Println("Go Web Assembly")
    js.Global().Set("sayHello", js.FuncOf(sayHello))
	<-make(chan bool)	// Stops program from exiting, 
						// makes go wait for this channel to exit 
						// (it never will)
}
// Use this to compile for WASM
// GOOS=js GOARCH=wasm go build -o asparser.wasm