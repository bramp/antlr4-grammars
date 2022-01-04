// Code generated from brainfuck.g4 by ANTLR 4.9.3. DO NOT EDIT.

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

// EnterFile_ is called when production file_ is entered.
func (s *BasebrainfuckListener) EnterFile_(ctx *File_Context) {}

// ExitFile_ is called when production file_ is exited.
func (s *BasebrainfuckListener) ExitFile_(ctx *File_Context) {}

// EnterStatement is called when production statement is entered.
func (s *BasebrainfuckListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasebrainfuckListener) ExitStatement(ctx *StatementContext) {}

// EnterOpcode is called when production opcode is entered.
func (s *BasebrainfuckListener) EnterOpcode(ctx *OpcodeContext) {}

// ExitOpcode is called when production opcode is exited.
func (s *BasebrainfuckListener) ExitOpcode(ctx *OpcodeContext) {}
