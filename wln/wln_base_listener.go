// Code generated from wln.g4 by ANTLR 4.9.3. DO NOT EDIT.

package wln // wln
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasewlnListener is a complete listener for a parse tree produced by wlnParser.
type BasewlnListener struct{}

var _ wlnListener = &BasewlnListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasewlnListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasewlnListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasewlnListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasewlnListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterWln is called when production wln is entered.
func (s *BasewlnListener) EnterWln(ctx *WlnContext) {}

// ExitWln is called when production wln is exited.
func (s *BasewlnListener) ExitWln(ctx *WlnContext) {}

// EnterGroup is called when production group is entered.
func (s *BasewlnListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BasewlnListener) ExitGroup(ctx *GroupContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BasewlnListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BasewlnListener) ExitSymbol(ctx *SymbolContext) {}
