// Generated from arithmetic.g4 by ANTLR 4.7.

package arithmetic // arithmetic
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasearithmeticVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasearithmeticVisitor) VisitEquation(ctx *EquationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearithmeticVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearithmeticVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearithmeticVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearithmeticVisitor) VisitSignedAtom(ctx *SignedAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearithmeticVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearithmeticVisitor) VisitScientific(ctx *ScientificContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearithmeticVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearithmeticVisitor) VisitRelop(ctx *RelopContext) interface{} {
	return v.VisitChildren(ctx)
}
