// Generated from redcode.g4 by ANTLR 4.7.

package redcode // redcode
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by redcodeParser.
type redcodeVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by redcodeParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by redcodeParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by redcodeParser#instruction.
	VisitInstruction(ctx *InstructionContext) interface{}

	// Visit a parse tree produced by redcodeParser#opcode.
	VisitOpcode(ctx *OpcodeContext) interface{}

	// Visit a parse tree produced by redcodeParser#modifier.
	VisitModifier(ctx *ModifierContext) interface{}

	// Visit a parse tree produced by redcodeParser#mmode.
	VisitMmode(ctx *MmodeContext) interface{}

	// Visit a parse tree produced by redcodeParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by redcodeParser#comment.
	VisitComment(ctx *CommentContext) interface{}
}
