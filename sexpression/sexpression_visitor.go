// Generated from sexpression.g4 by ANTLR 4.7.

package sexpression // sexpression
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by sexpressionParser.
type sexpressionVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by sexpressionParser#sexpr.
	VisitSexpr(ctx *SexprContext) interface{}

	// Visit a parse tree produced by sexpressionParser#item.
	VisitItem(ctx *ItemContext) interface{}

	// Visit a parse tree produced by sexpressionParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by sexpressionParser#atom.
	VisitAtom(ctx *AtomContext) interface{}
}
