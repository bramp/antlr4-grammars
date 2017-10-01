// Generated from RFilter.g4 by ANTLR 4.7.

package r // RFilter
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRFilterListener is a complete listener for a parse tree produced by RFilter.
type BaseRFilterListener struct{}

var _ RFilterListener = &BaseRFilterListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRFilterListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRFilterListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRFilterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRFilterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStream is called when production stream is entered.
func (s *BaseRFilterListener) EnterStream(ctx *StreamContext) {}

// ExitStream is called when production stream is exited.
func (s *BaseRFilterListener) ExitStream(ctx *StreamContext) {}

// EnterEat is called when production eat is entered.
func (s *BaseRFilterListener) EnterEat(ctx *EatContext) {}

// ExitEat is called when production eat is exited.
func (s *BaseRFilterListener) ExitEat(ctx *EatContext) {}

// EnterElem is called when production elem is entered.
func (s *BaseRFilterListener) EnterElem(ctx *ElemContext) {}

// ExitElem is called when production elem is exited.
func (s *BaseRFilterListener) ExitElem(ctx *ElemContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseRFilterListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseRFilterListener) ExitAtom(ctx *AtomContext) {}

// EnterOp is called when production op is entered.
func (s *BaseRFilterListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BaseRFilterListener) ExitOp(ctx *OpContext) {}
