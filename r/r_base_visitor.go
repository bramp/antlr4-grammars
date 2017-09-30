// Generated from R.g4 by ANTLR 4.7.

package r // R
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRVisitor) VisitExprlist(ctx *ExprlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRVisitor) VisitFormlist(ctx *FormlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRVisitor) VisitForm(ctx *FormContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRVisitor) VisitSublist(ctx *SublistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRVisitor) VisitSub(ctx *SubContext) interface{} {
	return v.VisitChildren(ctx)
}
