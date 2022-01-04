// Code generated from ctl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package ctl // ctl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasectlListener is a complete listener for a parse tree produced by ctlParser.
type BasectlListener struct{}

var _ ctlListener = &BasectlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasectlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasectlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasectlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasectlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProposition is called when production proposition is entered.
func (s *BasectlListener) EnterProposition(ctx *PropositionContext) {}

// ExitProposition is called when production proposition is exited.
func (s *BasectlListener) ExitProposition(ctx *PropositionContext) {}
