// Generated from tinyc.g4 by ANTLR 4.7.

package tinyc // tinyc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by tinycParser.
type tinycVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by tinycParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by tinycParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by tinycParser#paren_expr.
	VisitParen_expr(ctx *Paren_exprContext) interface{}

	// Visit a parse tree produced by tinycParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by tinycParser#test.
	VisitTest(ctx *TestContext) interface{}

	// Visit a parse tree produced by tinycParser#sum.
	VisitSum(ctx *SumContext) interface{}

	// Visit a parse tree produced by tinycParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by tinycParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by tinycParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}
}
