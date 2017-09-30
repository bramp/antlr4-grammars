// Generated from Abnf.g4 by ANTLR 4.7.

package abnf // Abnf
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseAbnfVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAbnfVisitor) VisitRulelist(ctx *RulelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbnfVisitor) VisitRule_(ctx *Rule_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbnfVisitor) VisitElements(ctx *ElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbnfVisitor) VisitAlternation(ctx *AlternationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbnfVisitor) VisitConcatenation(ctx *ConcatenationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbnfVisitor) VisitRepetition(ctx *RepetitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbnfVisitor) VisitRepeat(ctx *RepeatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbnfVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbnfVisitor) VisitGroup(ctx *GroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbnfVisitor) VisitOption(ctx *OptionContext) interface{} {
	return v.VisitChildren(ctx)
}
