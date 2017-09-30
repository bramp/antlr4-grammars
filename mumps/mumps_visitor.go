// Generated from mumps.g4 by ANTLR 4.7.

package mumps // mumps
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by mumpsParser.
type mumpsVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by mumpsParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by mumpsParser#eof.
	VisitEof(ctx *EofContext) interface{}

	// Visit a parse tree produced by mumpsParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by mumpsParser#code.
	VisitCode(ctx *CodeContext) interface{}

	// Visit a parse tree produced by mumpsParser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by mumpsParser#routinedecl.
	VisitRoutinedecl(ctx *RoutinedeclContext) interface{}

	// Visit a parse tree produced by mumpsParser#paramlist.
	VisitParamlist(ctx *ParamlistContext) interface{}

	// Visit a parse tree produced by mumpsParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by mumpsParser#subproc.
	VisitSubproc(ctx *SubprocContext) interface{}

	// Visit a parse tree produced by mumpsParser#command.
	VisitCommand(ctx *CommandContext) interface{}

	// Visit a parse tree produced by mumpsParser#postcondition.
	VisitPostcondition(ctx *PostconditionContext) interface{}

	// Visit a parse tree produced by mumpsParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by mumpsParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by mumpsParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by mumpsParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by mumpsParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by mumpsParser#break_.
	VisitBreak_(ctx *Break_Context) interface{}

	// Visit a parse tree produced by mumpsParser#do_.
	VisitDo_(ctx *Do_Context) interface{}

	// Visit a parse tree produced by mumpsParser#for_.
	VisitFor_(ctx *For_Context) interface{}

	// Visit a parse tree produced by mumpsParser#halt_.
	VisitHalt_(ctx *Halt_Context) interface{}

	// Visit a parse tree produced by mumpsParser#hang_.
	VisitHang_(ctx *Hang_Context) interface{}

	// Visit a parse tree produced by mumpsParser#if_.
	VisitIf_(ctx *If_Context) interface{}

	// Visit a parse tree produced by mumpsParser#kill_.
	VisitKill_(ctx *Kill_Context) interface{}

	// Visit a parse tree produced by mumpsParser#merge_.
	VisitMerge_(ctx *Merge_Context) interface{}

	// Visit a parse tree produced by mumpsParser#new_.
	VisitNew_(ctx *New_Context) interface{}

	// Visit a parse tree produced by mumpsParser#quit_.
	VisitQuit_(ctx *Quit_Context) interface{}

	// Visit a parse tree produced by mumpsParser#read_.
	VisitRead_(ctx *Read_Context) interface{}

	// Visit a parse tree produced by mumpsParser#set_.
	VisitSet_(ctx *Set_Context) interface{}

	// Visit a parse tree produced by mumpsParser#view_.
	VisitView_(ctx *View_Context) interface{}

	// Visit a parse tree produced by mumpsParser#write_.
	VisitWrite_(ctx *Write_Context) interface{}

	// Visit a parse tree produced by mumpsParser#xecute_.
	VisitXecute_(ctx *Xecute_Context) interface{}

	// Visit a parse tree produced by mumpsParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by mumpsParser#arglist.
	VisitArglist(ctx *ArglistContext) interface{}

	// Visit a parse tree produced by mumpsParser#arg.
	VisitArg(ctx *ArgContext) interface{}
}
