// Generated from brainfuck.g4 by ANTLR 4.7.

package brainfuck // brainfuck
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by brainfuckParser.
type brainfuckVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by brainfuckParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by brainfuckParser#opcode.
	VisitOpcode(ctx *OpcodeContext) interface{}
}
