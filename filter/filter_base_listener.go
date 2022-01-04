// Code generated from filter.g4 by ANTLR 4.9.3. DO NOT EDIT.

package filter // filter
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasefilterListener is a complete listener for a parse tree produced by filterParser.
type BasefilterListener struct{}

var _ filterListener = &BasefilterListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasefilterListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasefilterListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasefilterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasefilterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFilter_ is called when production filter_ is entered.
func (s *BasefilterListener) EnterFilter_(ctx *Filter_Context) {}

// ExitFilter_ is called when production filter_ is exited.
func (s *BasefilterListener) ExitFilter_(ctx *Filter_Context) {}

// EnterFiltercomp is called when production filtercomp is entered.
func (s *BasefilterListener) EnterFiltercomp(ctx *FiltercompContext) {}

// ExitFiltercomp is called when production filtercomp is exited.
func (s *BasefilterListener) ExitFiltercomp(ctx *FiltercompContext) {}

// EnterAnd_ is called when production and_ is entered.
func (s *BasefilterListener) EnterAnd_(ctx *And_Context) {}

// ExitAnd_ is called when production and_ is exited.
func (s *BasefilterListener) ExitAnd_(ctx *And_Context) {}

// EnterOr_ is called when production or_ is entered.
func (s *BasefilterListener) EnterOr_(ctx *Or_Context) {}

// ExitOr_ is called when production or_ is exited.
func (s *BasefilterListener) ExitOr_(ctx *Or_Context) {}

// EnterNot_ is called when production not_ is entered.
func (s *BasefilterListener) EnterNot_(ctx *Not_Context) {}

// ExitNot_ is called when production not_ is exited.
func (s *BasefilterListener) ExitNot_(ctx *Not_Context) {}

// EnterFilterlist is called when production filterlist is entered.
func (s *BasefilterListener) EnterFilterlist(ctx *FilterlistContext) {}

// ExitFilterlist is called when production filterlist is exited.
func (s *BasefilterListener) ExitFilterlist(ctx *FilterlistContext) {}

// EnterItem is called when production item is entered.
func (s *BasefilterListener) EnterItem(ctx *ItemContext) {}

// ExitItem is called when production item is exited.
func (s *BasefilterListener) ExitItem(ctx *ItemContext) {}

// EnterSimple is called when production simple is entered.
func (s *BasefilterListener) EnterSimple(ctx *SimpleContext) {}

// ExitSimple is called when production simple is exited.
func (s *BasefilterListener) ExitSimple(ctx *SimpleContext) {}

// EnterFiltertype is called when production filtertype is entered.
func (s *BasefilterListener) EnterFiltertype(ctx *FiltertypeContext) {}

// ExitFiltertype is called when production filtertype is exited.
func (s *BasefilterListener) ExitFiltertype(ctx *FiltertypeContext) {}

// EnterPresent is called when production present is entered.
func (s *BasefilterListener) EnterPresent(ctx *PresentContext) {}

// ExitPresent is called when production present is exited.
func (s *BasefilterListener) ExitPresent(ctx *PresentContext) {}

// EnterSubstring is called when production substring is entered.
func (s *BasefilterListener) EnterSubstring(ctx *SubstringContext) {}

// ExitSubstring is called when production substring is exited.
func (s *BasefilterListener) ExitSubstring(ctx *SubstringContext) {}

// EnterInitial is called when production initial is entered.
func (s *BasefilterListener) EnterInitial(ctx *InitialContext) {}

// ExitInitial is called when production initial is exited.
func (s *BasefilterListener) ExitInitial(ctx *InitialContext) {}

// EnterAny_ is called when production any_ is entered.
func (s *BasefilterListener) EnterAny_(ctx *Any_Context) {}

// ExitAny_ is called when production any_ is exited.
func (s *BasefilterListener) ExitAny_(ctx *Any_Context) {}

// EnterStarval is called when production starval is entered.
func (s *BasefilterListener) EnterStarval(ctx *StarvalContext) {}

// ExitStarval is called when production starval is exited.
func (s *BasefilterListener) ExitStarval(ctx *StarvalContext) {}

// EnterFinal_ is called when production final_ is entered.
func (s *BasefilterListener) EnterFinal_(ctx *Final_Context) {}

// ExitFinal_ is called when production final_ is exited.
func (s *BasefilterListener) ExitFinal_(ctx *Final_Context) {}

// EnterAttr is called when production attr is entered.
func (s *BasefilterListener) EnterAttr(ctx *AttrContext) {}

// ExitAttr is called when production attr is exited.
func (s *BasefilterListener) ExitAttr(ctx *AttrContext) {}

// EnterValue is called when production value is entered.
func (s *BasefilterListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasefilterListener) ExitValue(ctx *ValueContext) {}
