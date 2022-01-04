// Code generated from PMMN.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pmmn // PMMN
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePMMNListener is a complete listener for a parse tree produced by PMMNParser.
type BasePMMNListener struct{}

var _ PMMNListener = &BasePMMNListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePMMNListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePMMNListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePMMNListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePMMNListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCommandlist is called when production commandlist is entered.
func (s *BasePMMNListener) EnterCommandlist(ctx *CommandlistContext) {}

// ExitCommandlist is called when production commandlist is exited.
func (s *BasePMMNListener) ExitCommandlist(ctx *CommandlistContext) {}

// EnterBlock is called when production block is entered.
func (s *BasePMMNListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasePMMNListener) ExitBlock(ctx *BlockContext) {}

// EnterCommand is called when production command is entered.
func (s *BasePMMNListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BasePMMNListener) ExitCommand(ctx *CommandContext) {}

// EnterTest is called when production test is entered.
func (s *BasePMMNListener) EnterTest(ctx *TestContext) {}

// ExitTest is called when production test is exited.
func (s *BasePMMNListener) ExitTest(ctx *TestContext) {}

// EnterCounter is called when production counter is entered.
func (s *BasePMMNListener) EnterCounter(ctx *CounterContext) {}

// ExitCounter is called when production counter is exited.
func (s *BasePMMNListener) ExitCounter(ctx *CounterContext) {}
