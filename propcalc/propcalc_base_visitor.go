// Generated from propcalc.g4 by ANTLR 4.7.

package propcalc // propcalc
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasepropcalcVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasepropcalcVisitor) VisitProposition(ctx *PropositionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepropcalcVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepropcalcVisitor) VisitRelExpression(ctx *RelExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepropcalcVisitor) VisitAtoms(ctx *AtomsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepropcalcVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepropcalcVisitor) VisitEquiv(ctx *EquivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepropcalcVisitor) VisitImplies(ctx *ImpliesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepropcalcVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}
