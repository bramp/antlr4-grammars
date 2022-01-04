// Code generated from nanofuck.g4 by ANTLR 4.9.3. DO NOT EDIT.

package nanofuck // nanofuck
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasenanofuckListener is a complete listener for a parse tree produced by nanofuckParser.
type BasenanofuckListener struct{}

var _ nanofuckListener = &BasenanofuckListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasenanofuckListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasenanofuckListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasenanofuckListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasenanofuckListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExp is called when production exp is entered.
func (s *BasenanofuckListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BasenanofuckListener) ExitExp(ctx *ExpContext) {}
