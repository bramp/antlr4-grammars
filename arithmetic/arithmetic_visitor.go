// Generated from arithmetic.g4 by ANTLR 4.7.

package arithmetic // arithmetic
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by arithmeticParser.
type arithmeticVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by arithmeticParser#equation.
	VisitEquation(ctx *EquationContext) interface{}

	// Visit a parse tree produced by arithmeticParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by arithmeticParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by arithmeticParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by arithmeticParser#signedAtom.
	VisitSignedAtom(ctx *SignedAtomContext) interface{}

	// Visit a parse tree produced by arithmeticParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by arithmeticParser#scientific.
	VisitScientific(ctx *ScientificContext) interface{}

	// Visit a parse tree produced by arithmeticParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by arithmeticParser#relop.
	VisitRelop(ctx *RelopContext) interface{}
}
