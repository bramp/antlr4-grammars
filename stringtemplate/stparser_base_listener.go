// Generated from STParser.g4 by ANTLR 4.7.

package stringtemplate // STParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSTParserListener is a complete listener for a parse tree produced by STParser.
type BaseSTParserListener struct{}

var _ STParserListener = &BaseSTParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSTParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSTParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSTParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSTParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTemplate is called when production template is entered.
func (s *BaseSTParserListener) EnterTemplate(ctx *TemplateContext) {}

// ExitTemplate is called when production template is exited.
func (s *BaseSTParserListener) ExitTemplate(ctx *TemplateContext) {}

// EnterElements is called when production elements is entered.
func (s *BaseSTParserListener) EnterElements(ctx *ElementsContext) {}

// ExitElements is called when production elements is exited.
func (s *BaseSTParserListener) ExitElements(ctx *ElementsContext) {}

// EnterElement is called when production element is entered.
func (s *BaseSTParserListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseSTParserListener) ExitElement(ctx *ElementContext) {}

// EnterSingleElement is called when production singleElement is entered.
func (s *BaseSTParserListener) EnterSingleElement(ctx *SingleElementContext) {}

// ExitSingleElement is called when production singleElement is exited.
func (s *BaseSTParserListener) ExitSingleElement(ctx *SingleElementContext) {}

// EnterCompoundElement is called when production compoundElement is entered.
func (s *BaseSTParserListener) EnterCompoundElement(ctx *CompoundElementContext) {}

// ExitCompoundElement is called when production compoundElement is exited.
func (s *BaseSTParserListener) ExitCompoundElement(ctx *CompoundElementContext) {}

// EnterExprTag is called when production exprTag is entered.
func (s *BaseSTParserListener) EnterExprTag(ctx *ExprTagContext) {}

// ExitExprTag is called when production exprTag is exited.
func (s *BaseSTParserListener) ExitExprTag(ctx *ExprTagContext) {}

// EnterRegion is called when production region is entered.
func (s *BaseSTParserListener) EnterRegion(ctx *RegionContext) {}

// ExitRegion is called when production region is exited.
func (s *BaseSTParserListener) ExitRegion(ctx *RegionContext) {}

// EnterSubtemplate is called when production subtemplate is entered.
func (s *BaseSTParserListener) EnterSubtemplate(ctx *SubtemplateContext) {}

// ExitSubtemplate is called when production subtemplate is exited.
func (s *BaseSTParserListener) ExitSubtemplate(ctx *SubtemplateContext) {}

// EnterIfstat is called when production ifstat is entered.
func (s *BaseSTParserListener) EnterIfstat(ctx *IfstatContext) {}

// ExitIfstat is called when production ifstat is exited.
func (s *BaseSTParserListener) ExitIfstat(ctx *IfstatContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BaseSTParserListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BaseSTParserListener) ExitConditional(ctx *ConditionalContext) {}

// EnterAndConditional is called when production andConditional is entered.
func (s *BaseSTParserListener) EnterAndConditional(ctx *AndConditionalContext) {}

// ExitAndConditional is called when production andConditional is exited.
func (s *BaseSTParserListener) ExitAndConditional(ctx *AndConditionalContext) {}

// EnterNotConditional is called when production notConditional is entered.
func (s *BaseSTParserListener) EnterNotConditional(ctx *NotConditionalContext) {}

// ExitNotConditional is called when production notConditional is exited.
func (s *BaseSTParserListener) ExitNotConditional(ctx *NotConditionalContext) {}

// EnterNotConditionalExpr is called when production notConditionalExpr is entered.
func (s *BaseSTParserListener) EnterNotConditionalExpr(ctx *NotConditionalExprContext) {}

// ExitNotConditionalExpr is called when production notConditionalExpr is exited.
func (s *BaseSTParserListener) ExitNotConditionalExpr(ctx *NotConditionalExprContext) {}

// EnterExprOptions is called when production exprOptions is entered.
func (s *BaseSTParserListener) EnterExprOptions(ctx *ExprOptionsContext) {}

// ExitExprOptions is called when production exprOptions is exited.
func (s *BaseSTParserListener) ExitExprOptions(ctx *ExprOptionsContext) {}

// EnterOption is called when production option is entered.
func (s *BaseSTParserListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseSTParserListener) ExitOption(ctx *OptionContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSTParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSTParserListener) ExitExpr(ctx *ExprContext) {}

// EnterMapExpr is called when production mapExpr is entered.
func (s *BaseSTParserListener) EnterMapExpr(ctx *MapExprContext) {}

// ExitMapExpr is called when production mapExpr is exited.
func (s *BaseSTParserListener) ExitMapExpr(ctx *MapExprContext) {}

// EnterMemberExpr is called when production memberExpr is entered.
func (s *BaseSTParserListener) EnterMemberExpr(ctx *MemberExprContext) {}

// ExitMemberExpr is called when production memberExpr is exited.
func (s *BaseSTParserListener) ExitMemberExpr(ctx *MemberExprContext) {}

// EnterMapTemplateRef is called when production mapTemplateRef is entered.
func (s *BaseSTParserListener) EnterMapTemplateRef(ctx *MapTemplateRefContext) {}

// ExitMapTemplateRef is called when production mapTemplateRef is exited.
func (s *BaseSTParserListener) ExitMapTemplateRef(ctx *MapTemplateRefContext) {}

// EnterIncludeExpr is called when production includeExpr is entered.
func (s *BaseSTParserListener) EnterIncludeExpr(ctx *IncludeExprContext) {}

// ExitIncludeExpr is called when production includeExpr is exited.
func (s *BaseSTParserListener) ExitIncludeExpr(ctx *IncludeExprContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseSTParserListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseSTParserListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterList is called when production list is entered.
func (s *BaseSTParserListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseSTParserListener) ExitList(ctx *ListContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseSTParserListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseSTParserListener) ExitArgs(ctx *ArgsContext) {}

// EnterArgExprList is called when production argExprList is entered.
func (s *BaseSTParserListener) EnterArgExprList(ctx *ArgExprListContext) {}

// ExitArgExprList is called when production argExprList is exited.
func (s *BaseSTParserListener) ExitArgExprList(ctx *ArgExprListContext) {}

// EnterNamedArg is called when production namedArg is entered.
func (s *BaseSTParserListener) EnterNamedArg(ctx *NamedArgContext) {}

// ExitNamedArg is called when production namedArg is exited.
func (s *BaseSTParserListener) ExitNamedArg(ctx *NamedArgContext) {}
