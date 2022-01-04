// Code generated from ltl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package ltl // ltl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseltlListener is a complete listener for a parse tree produced by ltlParser.
type BaseltlListener struct{}

var _ ltlListener = &BaseltlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseltlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseltlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseltlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseltlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProposition is called when production proposition is entered.
func (s *BaseltlListener) EnterProposition(ctx *PropositionContext) {}

// ExitProposition is called when production proposition is exited.
func (s *BaseltlListener) ExitProposition(ctx *PropositionContext) {}
