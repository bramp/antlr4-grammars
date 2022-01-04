// Code generated from nanofuck.g4 by ANTLR 4.9.3. DO NOT EDIT.

package nanofuck // nanofuck
import "github.com/antlr/antlr4/runtime/Go/antlr"

// nanofuckListener is a complete listener for a parse tree produced by nanofuckParser.
type nanofuckListener interface {
	antlr.ParseTreeListener

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)
}
