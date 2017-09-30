// Generated from propcalc.g4 by ANTLR 4.7.

package propcalc // propcalc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by propcalcParser.
type propcalcVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by propcalcParser#proposition.
	VisitProposition(ctx *PropositionContext) interface{}

	// Visit a parse tree produced by propcalcParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by propcalcParser#relExpression.
	VisitRelExpression(ctx *RelExpressionContext) interface{}

	// Visit a parse tree produced by propcalcParser#atoms.
	VisitAtoms(ctx *AtomsContext) interface{}

	// Visit a parse tree produced by propcalcParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by propcalcParser#equiv.
	VisitEquiv(ctx *EquivContext) interface{}

	// Visit a parse tree produced by propcalcParser#implies.
	VisitImplies(ctx *ImpliesContext) interface{}

	// Visit a parse tree produced by propcalcParser#variable.
	VisitVariable(ctx *VariableContext) interface{}
}
