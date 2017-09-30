// Generated from tnt.g4 by ANTLR 4.7.

package tnt // tnt
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by tntParser.
type tntVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by tntParser#equation.
	VisitEquation(ctx *EquationContext) interface{}

	// Visit a parse tree produced by tntParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by tntParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by tntParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by tntParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by tntParser#forevery.
	VisitForevery(ctx *ForeveryContext) interface{}

	// Visit a parse tree produced by tntParser#exists.
	VisitExists(ctx *ExistsContext) interface{}
}
