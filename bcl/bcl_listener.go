// Code generated from bcl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bcl // bcl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// bclListener is a complete listener for a parse tree produced by bclParser.
type bclListener interface {
	antlr.ParseTreeListener

	// EnterBcl is called when entering the bcl production.
	EnterBcl(c *BclContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// ExitBcl is called when exiting the bcl production.
	ExitBcl(c *BclContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)
}
