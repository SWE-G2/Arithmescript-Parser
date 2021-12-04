# Arithmescript-Parser
A library for converting Arithmescript markup into LaTeX.

## Tests
To run the tests for the parser library...
```
git clone https://github.com/SWE-G2/Arithmescript-Parser
cd Arithmescript-Parser/parser
go install
go test
```

### Compile for WASM
```
make compile-wasm
```

### Misc

If your working with the JS bindings, add this to your VSCode workspace settings:
```json
"go.toolsEnvVars": {

    "GOARCH": "wasm",
    "GOOS": "js",
},
"go.testEnvVars": {
    "GOARCH": "wasm",
    "GOOS": "js",
}
```
