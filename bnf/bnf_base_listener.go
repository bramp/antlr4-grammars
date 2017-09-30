// Generated from bnf.g4 by ANTLR 4.7.

package bnf // bnf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasebnfListener is a complete listener for a parse tree produced by bnfParser.
type BasebnfListener struct{}

var _ bnfListener = &BasebnfListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasebnfListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasebnfListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasebnfListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasebnfListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRulelist is called when production rulelist is entered.
func (s *BasebnfListener) EnterRulelist(ctx *RulelistContext) {}

// ExitRulelist is called when production rulelist is exited.
func (s *BasebnfListener) ExitRulelist(ctx *RulelistContext) {}

// EnterRule_ is called when production rule_ is entered.
func (s *BasebnfListener) EnterRule_(ctx *Rule_Context) {}

// ExitRule_ is called when production rule_ is exited.
func (s *BasebnfListener) ExitRule_(ctx *Rule_Context) {}

// EnterLhs is called when production lhs is entered.
func (s *BasebnfListener) EnterLhs(ctx *LhsContext) {}

// ExitLhs is called when production lhs is exited.
func (s *BasebnfListener) ExitLhs(ctx *LhsContext) {}

// EnterRhs is called when production rhs is entered.
func (s *BasebnfListener) EnterRhs(ctx *RhsContext) {}

// ExitRhs is called when production rhs is exited.
func (s *BasebnfListener) ExitRhs(ctx *RhsContext) {}

// EnterAlternatives is called when production alternatives is entered.
func (s *BasebnfListener) EnterAlternatives(ctx *AlternativesContext) {}

// ExitAlternatives is called when production alternatives is exited.
func (s *BasebnfListener) ExitAlternatives(ctx *AlternativesContext) {}

// EnterAlternative is called when production alternative is entered.
func (s *BasebnfListener) EnterAlternative(ctx *AlternativeContext) {}

// ExitAlternative is called when production alternative is exited.
func (s *BasebnfListener) ExitAlternative(ctx *AlternativeContext) {}

// EnterElement is called when production element is entered.
func (s *BasebnfListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BasebnfListener) ExitElement(ctx *ElementContext) {}

// EnterOptional is called when production optional is entered.
func (s *BasebnfListener) EnterOptional(ctx *OptionalContext) {}

// ExitOptional is called when production optional is exited.
func (s *BasebnfListener) ExitOptional(ctx *OptionalContext) {}

// EnterZeroormore is called when production zeroormore is entered.
func (s *BasebnfListener) EnterZeroormore(ctx *ZeroormoreContext) {}

// ExitZeroormore is called when production zeroormore is exited.
func (s *BasebnfListener) ExitZeroormore(ctx *ZeroormoreContext) {}

// EnterOneormore is called when production oneormore is entered.
func (s *BasebnfListener) EnterOneormore(ctx *OneormoreContext) {}

// ExitOneormore is called when production oneormore is exited.
func (s *BasebnfListener) ExitOneormore(ctx *OneormoreContext) {}

// EnterText is called when production text is entered.
func (s *BasebnfListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BasebnfListener) ExitText(ctx *TextContext) {}

// EnterId is called when production id is entered.
func (s *BasebnfListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BasebnfListener) ExitId(ctx *IdContext) {}

// EnterRuleid is called when production ruleid is entered.
func (s *BasebnfListener) EnterRuleid(ctx *RuleidContext) {}

// ExitRuleid is called when production ruleid is exited.
func (s *BasebnfListener) ExitRuleid(ctx *RuleidContext) {}
