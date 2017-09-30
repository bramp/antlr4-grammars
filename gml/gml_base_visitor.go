// Generated from gml.g4 by ANTLR 4.7.

package gml // gml
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasegmlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasegmlVisitor) VisitGraph(ctx *GraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegmlVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegmlVisitor) VisitKv(ctx *KvContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegmlVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegmlVisitor) VisitKey(ctx *KeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegmlVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegmlVisitor) VisitRealnum(ctx *RealnumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegmlVisitor) VisitStr(ctx *StrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegmlVisitor) VisitStringliteral(ctx *StringliteralContext) interface{} {
	return v.VisitChildren(ctx)
}
