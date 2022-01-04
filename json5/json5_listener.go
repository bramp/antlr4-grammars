// Code generated from JSON5.g4 by ANTLR 4.9.3. DO NOT EDIT.

package json5 // JSON5
import "github.com/antlr/antlr4/runtime/Go/antlr"

// JSON5Listener is a complete listener for a parse tree produced by JSON5Parser.
type JSON5Listener interface {
	antlr.ParseTreeListener

	// EnterJson5 is called when entering the json5 production.
	EnterJson5(c *Json5Context)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterArr is called when entering the arr production.
	EnterArr(c *ArrContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitJson5 is called when exiting the json5 production.
	ExitJson5(c *Json5Context)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitArr is called when exiting the arr production.
	ExitArr(c *ArrContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
