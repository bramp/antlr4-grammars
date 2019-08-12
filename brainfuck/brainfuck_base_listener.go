// Code generated from brainfuck.g4 by ANTLR 4.7.2. DO NOT EDIT.

package brainfuck // brainfuck
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasebrainfuckListener is a complete listener for a parse tree produced by brainfuckParser.
type BasebrainfuckListener struct{}

var _ brainfuckListener = &BasebrainfuckListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasebrainfuckListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasebrainfuckListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasebrainfuckListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasebrainfuckListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BasebrainfuckListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BasebrainfuckListener) ExitFile(ctx *FileContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasebrainfuckListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasebrainfuckListener) ExitStatement(ctx *StatementContext) {}

// EnterOpcode is called when production opcode is entered.
func (s *BasebrainfuckListener) EnterOpcode(ctx *OpcodeContext) {}

// ExitOpcode is called when production opcode is exited.
func (s *BasebrainfuckListener) ExitOpcode(ctx *OpcodeContext) {}
