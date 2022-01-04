// Code generated from R.g4 by ANTLR 4.9.3. DO NOT EDIT.

package r // R
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRListener is a complete listener for a parse tree produced by RParser.
type BaseRListener struct{}

var _ RListener = &BaseRListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseRListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseRListener) ExitProg(ctx *ProgContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseRListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseRListener) ExitExpr(ctx *ExprContext) {}

// EnterExprlist is called when production exprlist is entered.
func (s *BaseRListener) EnterExprlist(ctx *ExprlistContext) {}

// ExitExprlist is called when production exprlist is exited.
func (s *BaseRListener) ExitExprlist(ctx *ExprlistContext) {}

// EnterFormlist is called when production formlist is entered.
func (s *BaseRListener) EnterFormlist(ctx *FormlistContext) {}

// ExitFormlist is called when production formlist is exited.
func (s *BaseRListener) ExitFormlist(ctx *FormlistContext) {}

// EnterForm is called when production form is entered.
func (s *BaseRListener) EnterForm(ctx *FormContext) {}

// ExitForm is called when production form is exited.
func (s *BaseRListener) ExitForm(ctx *FormContext) {}

// EnterSublist is called when production sublist is entered.
func (s *BaseRListener) EnterSublist(ctx *SublistContext) {}

// ExitSublist is called when production sublist is exited.
func (s *BaseRListener) ExitSublist(ctx *SublistContext) {}

// EnterSub is called when production sub is entered.
func (s *BaseRListener) EnterSub(ctx *SubContext) {}

// ExitSub is called when production sub is exited.
func (s *BaseRListener) ExitSub(ctx *SubContext) {}
