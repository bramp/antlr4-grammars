// Generated from bnf.g4 by ANTLR 4.7.

package bnf // bnf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by bnfParser.
type bnfVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by bnfParser#rulelist.
	VisitRulelist(ctx *RulelistContext) interface{}

	// Visit a parse tree produced by bnfParser#rule_.
	VisitRule_(ctx *Rule_Context) interface{}

	// Visit a parse tree produced by bnfParser#lhs.
	VisitLhs(ctx *LhsContext) interface{}

	// Visit a parse tree produced by bnfParser#rhs.
	VisitRhs(ctx *RhsContext) interface{}

	// Visit a parse tree produced by bnfParser#alternatives.
	VisitAlternatives(ctx *AlternativesContext) interface{}

	// Visit a parse tree produced by bnfParser#alternative.
	VisitAlternative(ctx *AlternativeContext) interface{}

	// Visit a parse tree produced by bnfParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by bnfParser#optional.
	VisitOptional(ctx *OptionalContext) interface{}

	// Visit a parse tree produced by bnfParser#zeroormore.
	VisitZeroormore(ctx *ZeroormoreContext) interface{}

	// Visit a parse tree produced by bnfParser#oneormore.
	VisitOneormore(ctx *OneormoreContext) interface{}

	// Visit a parse tree produced by bnfParser#text.
	VisitText(ctx *TextContext) interface{}

	// Visit a parse tree produced by bnfParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by bnfParser#ruleid.
	VisitRuleid(ctx *RuleidContext) interface{}
}
