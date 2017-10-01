// Generated from RFilter.g4 by ANTLR 4.7.

package r // RFilter
import "github.com/antlr/antlr4/runtime/Go/antlr"

// RFilterListener is a complete listener for a parse tree produced by RFilter.
type RFilterListener interface {
	antlr.ParseTreeListener

	// EnterStream is called when entering the stream production.
	EnterStream(c *StreamContext)

	// EnterEat is called when entering the eat production.
	EnterEat(c *EatContext)

	// EnterElem is called when entering the elem production.
	EnterElem(c *ElemContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// ExitStream is called when exiting the stream production.
	ExitStream(c *StreamContext)

	// ExitEat is called when exiting the eat production.
	ExitEat(c *EatContext)

	// ExitElem is called when exiting the elem production.
	ExitElem(c *ElemContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)
}
