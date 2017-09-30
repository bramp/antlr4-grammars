// Generated from agc.g4 by ANTLR 4.7.

package agc // agc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by agcParser.
type agcVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by agcParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by agcParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by agcParser#blank_line.
	VisitBlank_line(ctx *Blank_lineContext) interface{}

	// Visit a parse tree produced by agcParser#comment_line.
	VisitComment_line(ctx *Comment_lineContext) interface{}

	// Visit a parse tree produced by agcParser#instruction_line.
	VisitInstruction_line(ctx *Instruction_lineContext) interface{}

	// Visit a parse tree produced by agcParser#erase_line.
	VisitErase_line(ctx *Erase_lineContext) interface{}

	// Visit a parse tree produced by agcParser#assignment_line.
	VisitAssignment_line(ctx *Assignment_lineContext) interface{}

	// Visit a parse tree produced by agcParser#opcodes.
	VisitOpcodes(ctx *OpcodesContext) interface{}

	// Visit a parse tree produced by agcParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by agcParser#ws.
	VisitWs(ctx *WsContext) interface{}

	// Visit a parse tree produced by agcParser#eol.
	VisitEol(ctx *EolContext) interface{}

	// Visit a parse tree produced by agcParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by agcParser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by agcParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by agcParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by agcParser#multiplyingExpression.
	VisitMultiplyingExpression(ctx *MultiplyingExpressionContext) interface{}

	// Visit a parse tree produced by agcParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by agcParser#inte.
	VisitInte(ctx *InteContext) interface{}

	// Visit a parse tree produced by agcParser#decimal.
	VisitDecimal(ctx *DecimalContext) interface{}

	// Visit a parse tree produced by agcParser#register.
	VisitRegister(ctx *RegisterContext) interface{}

	// Visit a parse tree produced by agcParser#opcode.
	VisitOpcode(ctx *OpcodeContext) interface{}

	// Visit a parse tree produced by agcParser#axt_opcode.
	VisitAxt_opcode(ctx *Axt_opcodeContext) interface{}

	// Visit a parse tree produced by agcParser#pseudo_opcode.
	VisitPseudo_opcode(ctx *Pseudo_opcodeContext) interface{}

	// Visit a parse tree produced by agcParser#standard_opcode.
	VisitStandard_opcode(ctx *Standard_opcodeContext) interface{}
}
