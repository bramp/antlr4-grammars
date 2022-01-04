// Code generated from asm6502.g4 by ANTLR 4.9.3. DO NOT EDIT.

package asm6502 // asm6502
import "github.com/antlr/antlr4/runtime/Go/antlr"

// asm6502Listener is a complete listener for a parse tree produced by asm6502Parser.
type asm6502Listener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterAssemblerinstruction is called when entering the assemblerinstruction production.
	EnterAssemblerinstruction(c *AssemblerinstructionContext)

	// EnterAssembleropcode is called when entering the assembleropcode production.
	EnterAssembleropcode(c *AssembleropcodeContext)

	// EnterLbl is called when entering the lbl production.
	EnterLbl(c *LblContext)

	// EnterArgumentlist is called when entering the argumentlist production.
	EnterArgumentlist(c *ArgumentlistContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterPrefix_ is called when entering the prefix_ production.
	EnterPrefix_(c *Prefix_Context)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterOpcode is called when entering the opcode production.
	EnterOpcode(c *OpcodeContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitAssemblerinstruction is called when exiting the assemblerinstruction production.
	ExitAssemblerinstruction(c *AssemblerinstructionContext)

	// ExitAssembleropcode is called when exiting the assembleropcode production.
	ExitAssembleropcode(c *AssembleropcodeContext)

	// ExitLbl is called when exiting the lbl production.
	ExitLbl(c *LblContext)

	// ExitArgumentlist is called when exiting the argumentlist production.
	ExitArgumentlist(c *ArgumentlistContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitPrefix_ is called when exiting the prefix_ production.
	ExitPrefix_(c *Prefix_Context)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitOpcode is called when exiting the opcode production.
	ExitOpcode(c *OpcodeContext)
}
