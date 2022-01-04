// Code generated from ctl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package ctl // ctl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ctlListener is a complete listener for a parse tree produced by ctlParser.
type ctlListener interface {
	antlr.ParseTreeListener

	// EnterProposition is called when entering the proposition production.
	EnterProposition(c *PropositionContext)

	// ExitProposition is called when exiting the proposition production.
	ExitProposition(c *PropositionContext)
}
