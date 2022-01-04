// Code generated from brainfuck.g4 by ANTLR 4.9.3. DO NOT EDIT.

package brainfuck // brainfuck
import "github.com/antlr/antlr4/runtime/Go/antlr"

// brainfuckListener is a complete listener for a parse tree produced by brainfuckParser.
type brainfuckListener interface {
	antlr.ParseTreeListener

	// EnterFile_ is called when entering the file_ production.
	EnterFile_(c *File_Context)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterOpcode is called when entering the opcode production.
	EnterOpcode(c *OpcodeContext)

	// ExitFile_ is called when exiting the file_ production.
	ExitFile_(c *File_Context)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitOpcode is called when exiting the opcode production.
	ExitOpcode(c *OpcodeContext)
}
