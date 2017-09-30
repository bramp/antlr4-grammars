// Generated from R.g4 by ANTLR 4.7.

package r // R
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by RParser.
type RVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by RParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by RParser#exprlist.
	VisitExprlist(ctx *ExprlistContext) interface{}

	// Visit a parse tree produced by RParser#formlist.
	VisitFormlist(ctx *FormlistContext) interface{}

	// Visit a parse tree produced by RParser#form.
	VisitForm(ctx *FormContext) interface{}

	// Visit a parse tree produced by RParser#sublist.
	VisitSublist(ctx *SublistContext) interface{}

	// Visit a parse tree produced by RParser#sub.
	VisitSub(ctx *SubContext) interface{}
}
