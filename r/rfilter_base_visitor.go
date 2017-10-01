// Generated from RFilter.g4 by ANTLR 4.7.

package r // RFilter
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRFilterVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRFilterVisitor) VisitStream(ctx *StreamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRFilterVisitor) VisitEat(ctx *EatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRFilterVisitor) VisitElem(ctx *ElemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRFilterVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRFilterVisitor) VisitOp(ctx *OpContext) interface{} {
	return v.VisitChildren(ctx)
}
