// Generated from fol.g4 by ANTLR 4.7.

package fol // fol
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasefolVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasefolVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefolVisitor) VisitFormula(ctx *FormulaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefolVisitor) VisitDisjunction(ctx *DisjunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefolVisitor) VisitConjunction(ctx *ConjunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefolVisitor) VisitNegation(ctx *NegationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefolVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefolVisitor) VisitPredicateTuple(ctx *PredicateTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefolVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefolVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefolVisitor) VisitFunctionTuple(ctx *FunctionTupleContext) interface{} {
	return v.VisitChildren(ctx)
}
