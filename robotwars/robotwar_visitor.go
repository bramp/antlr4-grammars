// Generated from robotwar.g4 by ANTLR 4.7.

package robotwars // robotwar
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by robotwarParser.
type robotwarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by robotwarParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by robotwarParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by robotwarParser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by robotwarParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by robotwarParser#accumstatement.
	VisitAccumstatement(ctx *AccumstatementContext) interface{}

	// Visit a parse tree produced by robotwarParser#accumexpression.
	VisitAccumexpression(ctx *AccumexpressionContext) interface{}

	// Visit a parse tree produced by robotwarParser#gosubstatement.
	VisitGosubstatement(ctx *GosubstatementContext) interface{}

	// Visit a parse tree produced by robotwarParser#gotostatement.
	VisitGotostatement(ctx *GotostatementContext) interface{}

	// Visit a parse tree produced by robotwarParser#tostatement.
	VisitTostatement(ctx *TostatementContext) interface{}

	// Visit a parse tree produced by robotwarParser#endsubstatement.
	VisitEndsubstatement(ctx *EndsubstatementContext) interface{}

	// Visit a parse tree produced by robotwarParser#ifstatement.
	VisitIfstatement(ctx *IfstatementContext) interface{}

	// Visit a parse tree produced by robotwarParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by robotwarParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by robotwarParser#operation.
	VisitOperation(ctx *OperationContext) interface{}

	// Visit a parse tree produced by robotwarParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by robotwarParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by robotwarParser#register.
	VisitRegister(ctx *RegisterContext) interface{}

	// Visit a parse tree produced by robotwarParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by robotwarParser#comment.
	VisitComment(ctx *CommentContext) interface{}
}
