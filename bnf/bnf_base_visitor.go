// Generated from bnf.g4 by ANTLR 4.7.

package bnf // bnf
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasebnfVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasebnfVisitor) VisitRulelist(ctx *RulelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebnfVisitor) VisitRule_(ctx *Rule_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebnfVisitor) VisitLhs(ctx *LhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebnfVisitor) VisitRhs(ctx *RhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebnfVisitor) VisitAlternatives(ctx *AlternativesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebnfVisitor) VisitAlternative(ctx *AlternativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebnfVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebnfVisitor) VisitOptional(ctx *OptionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebnfVisitor) VisitZeroormore(ctx *ZeroormoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebnfVisitor) VisitOneormore(ctx *OneormoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebnfVisitor) VisitText(ctx *TextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebnfVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebnfVisitor) VisitRuleid(ctx *RuleidContext) interface{} {
	return v.VisitChildren(ctx)
}
