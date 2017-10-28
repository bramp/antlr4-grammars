package antlr4

import (
	"fmt"

	"bramp.net/antlr4/json"                    // The parser
	"github.com/antlr/antlr4/runtime/Go/antlr" // The antlr library
)

// exampleListener is an event-driven callback for the parser.
type exampleListener struct {
	*json.BaseJSONListener // https://godoc.org/bramp.net/antlr4/json#BaseJSONListener
}

func (l *exampleListener) EnterObj(ctx *json.ObjContext) {
	fmt.Printf("Object: %s\n", ctx.GetText())

}

func (l *exampleListener) EnterPair(ctx *json.PairContext) {
	fmt.Printf("Pair: %s\n", ctx.GetText())

}

func (l *exampleListener) EnterArray(ctx *json.ArrayContext) {
	fmt.Printf("Array: %s\n", ctx.GetText())

}

// Example shows how to use the JSON lexer and parser.
func Example() {
	// Setup the input
	is := antlr.NewInputStream(`{"example": "json", "with": ["an", "array"]}`)

	// Create the Lexer
	lexer := json.NewJSONLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := json.NewJSONParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Finally walk the tree
	tree := p.Json()
	antlr.ParseTreeWalkerDefault.Walk(&exampleListener{}, tree)

	// Output:
	// Object: {"example":"json","with":["an","array"]}
	// Pair: "example":"json"
	// Pair: "with":["an","array"]
	// Array: ["an","array"]
}
