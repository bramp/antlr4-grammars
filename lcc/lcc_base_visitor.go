// Generated from lcc.g4 by ANTLR 4.7.

package lcc // lcc
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaselccVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaselccVisitor) VisitLcc(ctx *LccContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaselccVisitor) VisitTopic(ctx *TopicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaselccVisitor) VisitSubtopic(ctx *SubtopicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaselccVisitor) VisitSubclasses(ctx *SubclassesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaselccVisitor) VisitSubclass(ctx *SubclassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaselccVisitor) VisitCutters(ctx *CuttersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaselccVisitor) VisitCutter(ctx *CutterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaselccVisitor) VisitDate(ctx *DateContext) interface{} {
	return v.VisitChildren(ctx)
}
