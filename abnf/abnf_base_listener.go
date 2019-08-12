// Code generated from Abnf.g4 by ANTLR 4.7.2. DO NOT EDIT.

package abnf // Abnf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAbnfListener is a complete listener for a parse tree produced by AbnfParser.
type BaseAbnfListener struct{}

var _ AbnfListener = &BaseAbnfListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAbnfListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAbnfListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAbnfListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAbnfListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRulelist is called when production rulelist is entered.
func (s *BaseAbnfListener) EnterRulelist(ctx *RulelistContext) {}

// ExitRulelist is called when production rulelist is exited.
func (s *BaseAbnfListener) ExitRulelist(ctx *RulelistContext) {}

// EnterRule_ is called when production rule_ is entered.
func (s *BaseAbnfListener) EnterRule_(ctx *Rule_Context) {}

// ExitRule_ is called when production rule_ is exited.
func (s *BaseAbnfListener) ExitRule_(ctx *Rule_Context) {}

// EnterElements is called when production elements is entered.
func (s *BaseAbnfListener) EnterElements(ctx *ElementsContext) {}

// ExitElements is called when production elements is exited.
func (s *BaseAbnfListener) ExitElements(ctx *ElementsContext) {}

// EnterAlternation is called when production alternation is entered.
func (s *BaseAbnfListener) EnterAlternation(ctx *AlternationContext) {}

// ExitAlternation is called when production alternation is exited.
func (s *BaseAbnfListener) ExitAlternation(ctx *AlternationContext) {}

// EnterConcatenation is called when production concatenation is entered.
func (s *BaseAbnfListener) EnterConcatenation(ctx *ConcatenationContext) {}

// ExitConcatenation is called when production concatenation is exited.
func (s *BaseAbnfListener) ExitConcatenation(ctx *ConcatenationContext) {}

// EnterRepetition is called when production repetition is entered.
func (s *BaseAbnfListener) EnterRepetition(ctx *RepetitionContext) {}

// ExitRepetition is called when production repetition is exited.
func (s *BaseAbnfListener) ExitRepetition(ctx *RepetitionContext) {}

// EnterRepeat is called when production repeat is entered.
func (s *BaseAbnfListener) EnterRepeat(ctx *RepeatContext) {}

// ExitRepeat is called when production repeat is exited.
func (s *BaseAbnfListener) ExitRepeat(ctx *RepeatContext) {}

// EnterElement is called when production element is entered.
func (s *BaseAbnfListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseAbnfListener) ExitElement(ctx *ElementContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseAbnfListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseAbnfListener) ExitGroup(ctx *GroupContext) {}

// EnterOption is called when production option is entered.
func (s *BaseAbnfListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseAbnfListener) ExitOption(ctx *OptionContext) {}
