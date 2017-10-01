// Generated from STParser.g4 by ANTLR 4.7.

package stringtemplate // STParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// STParserListener is a complete listener for a parse tree produced by STParser.
type STParserListener interface {
	antlr.ParseTreeListener

	// EnterTemplate is called when entering the template production.
	EnterTemplate(c *TemplateContext)

	// EnterElements is called when entering the elements production.
	EnterElements(c *ElementsContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterSingleElement is called when entering the singleElement production.
	EnterSingleElement(c *SingleElementContext)

	// EnterCompoundElement is called when entering the compoundElement production.
	EnterCompoundElement(c *CompoundElementContext)

	// EnterExprTag is called when entering the exprTag production.
	EnterExprTag(c *ExprTagContext)

	// EnterRegion is called when entering the region production.
	EnterRegion(c *RegionContext)

	// EnterSubtemplate is called when entering the subtemplate production.
	EnterSubtemplate(c *SubtemplateContext)

	// EnterIfstat is called when entering the ifstat production.
	EnterIfstat(c *IfstatContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterAndConditional is called when entering the andConditional production.
	EnterAndConditional(c *AndConditionalContext)

	// EnterNotConditional is called when entering the notConditional production.
	EnterNotConditional(c *NotConditionalContext)

	// EnterNotConditionalExpr is called when entering the notConditionalExpr production.
	EnterNotConditionalExpr(c *NotConditionalExprContext)

	// EnterExprOptions is called when entering the exprOptions production.
	EnterExprOptions(c *ExprOptionsContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterMapExpr is called when entering the mapExpr production.
	EnterMapExpr(c *MapExprContext)

	// EnterMemberExpr is called when entering the memberExpr production.
	EnterMemberExpr(c *MemberExprContext)

	// EnterMapTemplateRef is called when entering the mapTemplateRef production.
	EnterMapTemplateRef(c *MapTemplateRefContext)

	// EnterIncludeExpr is called when entering the includeExpr production.
	EnterIncludeExpr(c *IncludeExprContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterArgExprList is called when entering the argExprList production.
	EnterArgExprList(c *ArgExprListContext)

	// EnterNamedArg is called when entering the namedArg production.
	EnterNamedArg(c *NamedArgContext)

	// ExitTemplate is called when exiting the template production.
	ExitTemplate(c *TemplateContext)

	// ExitElements is called when exiting the elements production.
	ExitElements(c *ElementsContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitSingleElement is called when exiting the singleElement production.
	ExitSingleElement(c *SingleElementContext)

	// ExitCompoundElement is called when exiting the compoundElement production.
	ExitCompoundElement(c *CompoundElementContext)

	// ExitExprTag is called when exiting the exprTag production.
	ExitExprTag(c *ExprTagContext)

	// ExitRegion is called when exiting the region production.
	ExitRegion(c *RegionContext)

	// ExitSubtemplate is called when exiting the subtemplate production.
	ExitSubtemplate(c *SubtemplateContext)

	// ExitIfstat is called when exiting the ifstat production.
	ExitIfstat(c *IfstatContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitAndConditional is called when exiting the andConditional production.
	ExitAndConditional(c *AndConditionalContext)

	// ExitNotConditional is called when exiting the notConditional production.
	ExitNotConditional(c *NotConditionalContext)

	// ExitNotConditionalExpr is called when exiting the notConditionalExpr production.
	ExitNotConditionalExpr(c *NotConditionalExprContext)

	// ExitExprOptions is called when exiting the exprOptions production.
	ExitExprOptions(c *ExprOptionsContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitMapExpr is called when exiting the mapExpr production.
	ExitMapExpr(c *MapExprContext)

	// ExitMemberExpr is called when exiting the memberExpr production.
	ExitMemberExpr(c *MemberExprContext)

	// ExitMapTemplateRef is called when exiting the mapTemplateRef production.
	ExitMapTemplateRef(c *MapTemplateRefContext)

	// ExitIncludeExpr is called when exiting the includeExpr production.
	ExitIncludeExpr(c *IncludeExprContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitArgExprList is called when exiting the argExprList production.
	ExitArgExprList(c *ArgExprListContext)

	// ExitNamedArg is called when exiting the namedArg production.
	ExitNamedArg(c *NamedArgContext)
}
