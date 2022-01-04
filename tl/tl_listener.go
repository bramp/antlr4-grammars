// Code generated from tl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package tl // tl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// tlListener is a complete listener for a parse tree produced by tlParser.
type tlListener interface {
	antlr.ParseTreeListener

	// EnterProposition is called when entering the proposition production.
	EnterProposition(c *PropositionContext)

	// ExitProposition is called when exiting the proposition production.
	ExitProposition(c *PropositionContext)
}
