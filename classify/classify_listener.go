// Code generated from classify.g4 by ANTLR 4.9.3. DO NOT EDIT.

package classify // classify
import "github.com/antlr/antlr4/runtime/Go/antlr"

// classifyListener is a complete listener for a parse tree produced by classifyParser.
type classifyListener interface {
	antlr.ParseTreeListener

	// EnterCodepoint is called when entering the codepoint production.
	EnterCodepoint(c *CodepointContext)

	// ExitCodepoint is called when exiting the codepoint production.
	ExitCodepoint(c *CodepointContext)
}
