// Code generated from redcode.g4 by ANTLR 4.9.3. DO NOT EDIT.

package redcode // redcode
import "github.com/antlr/antlr4/runtime/Go/antlr"

// redcodeListener is a complete listener for a parse tree produced by redcodeParser.
type redcodeListener interface {
	antlr.ParseTreeListener

	// EnterFile_ is called when entering the file_ production.
	EnterFile_(c *File_Context)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterOpcode is called when entering the opcode production.
	EnterOpcode(c *OpcodeContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterMmode is called when entering the mmode production.
	EnterMmode(c *MmodeContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// ExitFile_ is called when exiting the file_ production.
	ExitFile_(c *File_Context)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitOpcode is called when exiting the opcode production.
	ExitOpcode(c *OpcodeContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitMmode is called when exiting the mmode production.
	ExitMmode(c *MmodeContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)
}
