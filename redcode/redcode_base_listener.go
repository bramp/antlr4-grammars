// Code generated from redcode.g4 by ANTLR 4.7.2. DO NOT EDIT.

package redcode // redcode
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseredcodeListener is a complete listener for a parse tree produced by redcodeParser.
type BaseredcodeListener struct{}

var _ redcodeListener = &BaseredcodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseredcodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseredcodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseredcodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseredcodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseredcodeListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseredcodeListener) ExitFile(ctx *FileContext) {}

// EnterLine is called when production line is entered.
func (s *BaseredcodeListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseredcodeListener) ExitLine(ctx *LineContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseredcodeListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseredcodeListener) ExitInstruction(ctx *InstructionContext) {}

// EnterOpcode is called when production opcode is entered.
func (s *BaseredcodeListener) EnterOpcode(ctx *OpcodeContext) {}

// ExitOpcode is called when production opcode is exited.
func (s *BaseredcodeListener) ExitOpcode(ctx *OpcodeContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseredcodeListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseredcodeListener) ExitModifier(ctx *ModifierContext) {}

// EnterMmode is called when production mmode is entered.
func (s *BaseredcodeListener) EnterMmode(ctx *MmodeContext) {}

// ExitMmode is called when production mmode is exited.
func (s *BaseredcodeListener) ExitMmode(ctx *MmodeContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseredcodeListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseredcodeListener) ExitNumber(ctx *NumberContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseredcodeListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseredcodeListener) ExitComment(ctx *CommentContext) {}
