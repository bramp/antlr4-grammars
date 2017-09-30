// Generated from tnt.g4 by ANTLR 4.7.

package tnt // tnt
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasetntVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasetntVisitor) VisitEquation(ctx *EquationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetntVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetntVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetntVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetntVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetntVisitor) VisitForevery(ctx *ForeveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetntVisitor) VisitExists(ctx *ExistsContext) interface{} {
	return v.VisitChildren(ctx)
}
