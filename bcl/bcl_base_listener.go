// Code generated from bcl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bcl // bcl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasebclListener is a complete listener for a parse tree produced by bclParser.
type BasebclListener struct{}

var _ bclListener = &BasebclListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasebclListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasebclListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasebclListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasebclListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterBcl is called when production bcl is entered.
func (s *BasebclListener) EnterBcl(ctx *BclContext) {}

// ExitBcl is called when production bcl is exited.
func (s *BasebclListener) ExitBcl(ctx *BclContext) {}

// EnterTerm is called when production term is entered.
func (s *BasebclListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasebclListener) ExitTerm(ctx *TermContext) {}
