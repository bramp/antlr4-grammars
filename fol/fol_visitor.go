// Generated from fol.g4 by ANTLR 4.7.

package fol // fol
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by folParser.
type folVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by folParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by folParser#formula.
	VisitFormula(ctx *FormulaContext) interface{}

	// Visit a parse tree produced by folParser#disjunction.
	VisitDisjunction(ctx *DisjunctionContext) interface{}

	// Visit a parse tree produced by folParser#conjunction.
	VisitConjunction(ctx *ConjunctionContext) interface{}

	// Visit a parse tree produced by folParser#negation.
	VisitNegation(ctx *NegationContext) interface{}

	// Visit a parse tree produced by folParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by folParser#predicateTuple.
	VisitPredicateTuple(ctx *PredicateTupleContext) interface{}

	// Visit a parse tree produced by folParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by folParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by folParser#functionTuple.
	VisitFunctionTuple(ctx *FunctionTupleContext) interface{}
}
