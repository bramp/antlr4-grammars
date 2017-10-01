// Generated from STParser.g4 by ANTLR 4.7.

package stringtemplate // STParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSTParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSTParserVisitor) VisitTemplate(ctx *TemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitElements(ctx *ElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitSingleElement(ctx *SingleElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitCompoundElement(ctx *CompoundElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitExprTag(ctx *ExprTagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitRegion(ctx *RegionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitSubtemplate(ctx *SubtemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitIfstat(ctx *IfstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitConditional(ctx *ConditionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitAndConditional(ctx *AndConditionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitNotConditional(ctx *NotConditionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitNotConditionalExpr(ctx *NotConditionalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitExprOptions(ctx *ExprOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitOption(ctx *OptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitMapExpr(ctx *MapExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitMemberExpr(ctx *MemberExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitMapTemplateRef(ctx *MapTemplateRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitIncludeExpr(ctx *IncludeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitArgs(ctx *ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitArgExprList(ctx *ArgExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTParserVisitor) VisitNamedArg(ctx *NamedArgContext) interface{} {
	return v.VisitChildren(ctx)
}
