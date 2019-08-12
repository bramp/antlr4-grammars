// Code generated from gml.g4 by ANTLR 4.7.2. DO NOT EDIT.

package gml // gml
import "github.com/antlr/antlr4/runtime/Go/antlr"

// gmlListener is a complete listener for a parse tree produced by gmlParser.
type gmlListener interface {
	antlr.ParseTreeListener

	// EnterGraph is called when entering the graph production.
	EnterGraph(c *GraphContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterKv is called when entering the kv production.
	EnterKv(c *KvContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterRealnum is called when entering the realnum production.
	EnterRealnum(c *RealnumContext)

	// EnterStr is called when entering the str production.
	EnterStr(c *StrContext)

	// EnterStringliteral is called when entering the stringliteral production.
	EnterStringliteral(c *StringliteralContext)

	// ExitGraph is called when exiting the graph production.
	ExitGraph(c *GraphContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitKv is called when exiting the kv production.
	ExitKv(c *KvContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitRealnum is called when exiting the realnum production.
	ExitRealnum(c *RealnumContext)

	// ExitStr is called when exiting the str production.
	ExitStr(c *StrContext)

	// ExitStringliteral is called when exiting the stringliteral production.
	ExitStringliteral(c *StringliteralContext)
}
