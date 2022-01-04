// Code generated from sexpression.g4 by ANTLR 4.9.3. DO NOT EDIT.

package sexpression // sexpression
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesexpressionListener is a complete listener for a parse tree produced by sexpressionParser.
type BasesexpressionListener struct{}

var _ sexpressionListener = &BasesexpressionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesexpressionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesexpressionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesexpressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesexpressionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSexpr is called when production sexpr is entered.
func (s *BasesexpressionListener) EnterSexpr(ctx *SexprContext) {}

// ExitSexpr is called when production sexpr is exited.
func (s *BasesexpressionListener) ExitSexpr(ctx *SexprContext) {}

// EnterItem is called when production item is entered.
func (s *BasesexpressionListener) EnterItem(ctx *ItemContext) {}

// ExitItem is called when production item is exited.
func (s *BasesexpressionListener) ExitItem(ctx *ItemContext) {}

// EnterList_ is called when production list_ is entered.
func (s *BasesexpressionListener) EnterList_(ctx *List_Context) {}

// ExitList_ is called when production list_ is exited.
func (s *BasesexpressionListener) ExitList_(ctx *List_Context) {}

// EnterAtom is called when production atom is entered.
func (s *BasesexpressionListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasesexpressionListener) ExitAtom(ctx *AtomContext) {}
