// Generated from sexpression.g4 by ANTLR 4.7.

package sexpression // sexpression
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasesexpressionVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasesexpressionVisitor) VisitSexpr(ctx *SexprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesexpressionVisitor) VisitItem(ctx *ItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesexpressionVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesexpressionVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}
