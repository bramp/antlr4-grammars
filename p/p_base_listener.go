// Code generated from p.g4 by ANTLR 4.7.2. DO NOT EDIT.

package p // p
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasepListener is a complete listener for a parse tree produced by pParser.
type BasepListener struct{}

var _ pListener = &BasepListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasepListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasepListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasepListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasepListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasepListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasepListener) ExitProg(ctx *ProgContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BasepListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BasepListener) ExitSymbol(ctx *SymbolContext) {}

// EnterIterate is called when production iterate is entered.
func (s *BasepListener) EnterIterate(ctx *IterateContext) {}

// ExitIterate is called when production iterate is exited.
func (s *BasepListener) ExitIterate(ctx *IterateContext) {}

// EnterAtom is called when production atom is entered.
func (s *BasepListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasepListener) ExitAtom(ctx *AtomContext) {}
