// Code generated from PMMN.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pmmn // PMMN
import "github.com/antlr/antlr4/runtime/Go/antlr"

// PMMNListener is a complete listener for a parse tree produced by PMMNParser.
type PMMNListener interface {
	antlr.ParseTreeListener

	// EnterCommandlist is called when entering the commandlist production.
	EnterCommandlist(c *CommandlistContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterTest is called when entering the test production.
	EnterTest(c *TestContext)

	// EnterCounter is called when entering the counter production.
	EnterCounter(c *CounterContext)

	// ExitCommandlist is called when exiting the commandlist production.
	ExitCommandlist(c *CommandlistContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitTest is called when exiting the test production.
	ExitTest(c *TestContext)

	// ExitCounter is called when exiting the counter production.
	ExitCounter(c *CounterContext)
}
