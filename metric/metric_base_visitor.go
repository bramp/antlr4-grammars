// Generated from metric.g4 by ANTLR 4.7.

package metric // metric
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasemetricVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasemetricVisitor) VisitUom(ctx *UomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemetricVisitor) VisitMeasure(ctx *MeasureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemetricVisitor) VisitExponent(ctx *ExponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemetricVisitor) VisitPrefix(ctx *PrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemetricVisitor) VisitUnit(ctx *UnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemetricVisitor) VisitBaseunit(ctx *BaseunitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemetricVisitor) VisitDerivedunit(ctx *DerivedunitContext) interface{} {
	return v.VisitChildren(ctx)
}
