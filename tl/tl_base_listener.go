// Code generated from tl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package tl // tl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasetlListener is a complete listener for a parse tree produced by tlParser.
type BasetlListener struct{}

var _ tlListener = &BasetlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProposition is called when production proposition is entered.
func (s *BasetlListener) EnterProposition(ctx *PropositionContext) {}

// ExitProposition is called when production proposition is exited.
func (s *BasetlListener) ExitProposition(ctx *PropositionContext) {}
