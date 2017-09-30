// Generated from rpn.g4 by ANTLR 4.7.

package rpn // rpn
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaserpnVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaserpnVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserpnVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserpnVisitor) VisitOper(ctx *OperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserpnVisitor) VisitSignedAtom(ctx *SignedAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserpnVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserpnVisitor) VisitScientific(ctx *ScientificContext) interface{} {
	return v.VisitChildren(ctx)
}
