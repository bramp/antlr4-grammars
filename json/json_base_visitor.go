// Generated from JSON.g4 by ANTLR 4.7.

package json // JSON
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseJSONVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJSONVisitor) VisitJson(ctx *JsonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitObj(ctx *ObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}
