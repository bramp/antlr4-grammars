// Generated from fol.g4 by ANTLR 4.7.

package fol // fol
import "github.com/antlr/antlr4/runtime/Go/antlr"

// folListener is a complete listener for a parse tree produced by folParser.
type folListener interface {
	antlr.ParseTreeListener

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterFormula is called when entering the formula production.
	EnterFormula(c *FormulaContext)

	// EnterDisjunction is called when entering the disjunction production.
	EnterDisjunction(c *DisjunctionContext)

	// EnterConjunction is called when entering the conjunction production.
	EnterConjunction(c *ConjunctionContext)

	// EnterNegation is called when entering the negation production.
	EnterNegation(c *NegationContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterPredicateTuple is called when entering the predicateTuple production.
	EnterPredicateTuple(c *PredicateTupleContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFunctionTuple is called when entering the functionTuple production.
	EnterFunctionTuple(c *FunctionTupleContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitFormula is called when exiting the formula production.
	ExitFormula(c *FormulaContext)

	// ExitDisjunction is called when exiting the disjunction production.
	ExitDisjunction(c *DisjunctionContext)

	// ExitConjunction is called when exiting the conjunction production.
	ExitConjunction(c *ConjunctionContext)

	// ExitNegation is called when exiting the negation production.
	ExitNegation(c *NegationContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitPredicateTuple is called when exiting the predicateTuple production.
	ExitPredicateTuple(c *PredicateTupleContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFunctionTuple is called when exiting the functionTuple production.
	ExitFunctionTuple(c *FunctionTupleContext)
}
