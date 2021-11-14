compile-wasm:
	GOOS=js GOARCH=wasm go build -o bin/asparser.wasm bindings/js/js.go 