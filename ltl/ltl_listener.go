// Code generated from ltl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package ltl // ltl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ltlListener is a complete listener for a parse tree produced by ltlParser.
type ltlListener interface {
	antlr.ParseTreeListener

	// EnterProposition is called when entering the proposition production.
	EnterProposition(c *PropositionContext)

	// ExitProposition is called when exiting the proposition production.
	ExitProposition(c *PropositionContext)
}
