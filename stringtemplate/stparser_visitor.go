// Generated from STParser.g4 by ANTLR 4.7.

package stringtemplate // STParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by STParser.
type STParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by STParser#template.
	VisitTemplate(ctx *TemplateContext) interface{}

	// Visit a parse tree produced by STParser#elements.
	VisitElements(ctx *ElementsContext) interface{}

	// Visit a parse tree produced by STParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by STParser#singleElement.
	VisitSingleElement(ctx *SingleElementContext) interface{}

	// Visit a parse tree produced by STParser#compoundElement.
	VisitCompoundElement(ctx *CompoundElementContext) interface{}

	// Visit a parse tree produced by STParser#exprTag.
	VisitExprTag(ctx *ExprTagContext) interface{}

	// Visit a parse tree produced by STParser#region.
	VisitRegion(ctx *RegionContext) interface{}

	// Visit a parse tree produced by STParser#subtemplate.
	VisitSubtemplate(ctx *SubtemplateContext) interface{}

	// Visit a parse tree produced by STParser#ifstat.
	VisitIfstat(ctx *IfstatContext) interface{}

	// Visit a parse tree produced by STParser#conditional.
	VisitConditional(ctx *ConditionalContext) interface{}

	// Visit a parse tree produced by STParser#andConditional.
	VisitAndConditional(ctx *AndConditionalContext) interface{}

	// Visit a parse tree produced by STParser#notConditional.
	VisitNotConditional(ctx *NotConditionalContext) interface{}

	// Visit a parse tree produced by STParser#notConditionalExpr.
	VisitNotConditionalExpr(ctx *NotConditionalExprContext) interface{}

	// Visit a parse tree produced by STParser#exprOptions.
	VisitExprOptions(ctx *ExprOptionsContext) interface{}

	// Visit a parse tree produced by STParser#option.
	VisitOption(ctx *OptionContext) interface{}

	// Visit a parse tree produced by STParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by STParser#mapExpr.
	VisitMapExpr(ctx *MapExprContext) interface{}

	// Visit a parse tree produced by STParser#memberExpr.
	VisitMemberExpr(ctx *MemberExprContext) interface{}

	// Visit a parse tree produced by STParser#mapTemplateRef.
	VisitMapTemplateRef(ctx *MapTemplateRefContext) interface{}

	// Visit a parse tree produced by STParser#includeExpr.
	VisitIncludeExpr(ctx *IncludeExprContext) interface{}

	// Visit a parse tree produced by STParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by STParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by STParser#args.
	VisitArgs(ctx *ArgsContext) interface{}

	// Visit a parse tree produced by STParser#argExprList.
	VisitArgExprList(ctx *ArgExprListContext) interface{}

	// Visit a parse tree produced by STParser#namedArg.
	VisitNamedArg(ctx *NamedArgContext) interface{}
}
