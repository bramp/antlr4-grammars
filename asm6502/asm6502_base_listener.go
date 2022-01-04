// Code generated from asm6502.g4 by ANTLR 4.9.3. DO NOT EDIT.

package asm6502 // asm6502
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Baseasm6502Listener is a complete listener for a parse tree produced by asm6502Parser.
type Baseasm6502Listener struct{}

var _ asm6502Listener = &Baseasm6502Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Baseasm6502Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Baseasm6502Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Baseasm6502Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Baseasm6502Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *Baseasm6502Listener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *Baseasm6502Listener) ExitProg(ctx *ProgContext) {}

// EnterLine is called when production line is entered.
func (s *Baseasm6502Listener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *Baseasm6502Listener) ExitLine(ctx *LineContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *Baseasm6502Listener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *Baseasm6502Listener) ExitInstruction(ctx *InstructionContext) {}

// EnterAssemblerinstruction is called when production assemblerinstruction is entered.
func (s *Baseasm6502Listener) EnterAssemblerinstruction(ctx *AssemblerinstructionContext) {}

// ExitAssemblerinstruction is called when production assemblerinstruction is exited.
func (s *Baseasm6502Listener) ExitAssemblerinstruction(ctx *AssemblerinstructionContext) {}

// EnterAssembleropcode is called when production assembleropcode is entered.
func (s *Baseasm6502Listener) EnterAssembleropcode(ctx *AssembleropcodeContext) {}

// ExitAssembleropcode is called when production assembleropcode is exited.
func (s *Baseasm6502Listener) ExitAssembleropcode(ctx *AssembleropcodeContext) {}

// EnterLbl is called when production lbl is entered.
func (s *Baseasm6502Listener) EnterLbl(ctx *LblContext) {}

// ExitLbl is called when production lbl is exited.
func (s *Baseasm6502Listener) ExitLbl(ctx *LblContext) {}

// EnterArgumentlist is called when production argumentlist is entered.
func (s *Baseasm6502Listener) EnterArgumentlist(ctx *ArgumentlistContext) {}

// ExitArgumentlist is called when production argumentlist is exited.
func (s *Baseasm6502Listener) ExitArgumentlist(ctx *ArgumentlistContext) {}

// EnterLabel is called when production label is entered.
func (s *Baseasm6502Listener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *Baseasm6502Listener) ExitLabel(ctx *LabelContext) {}

// EnterArgument is called when production argument is entered.
func (s *Baseasm6502Listener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *Baseasm6502Listener) ExitArgument(ctx *ArgumentContext) {}

// EnterPrefix_ is called when production prefix_ is entered.
func (s *Baseasm6502Listener) EnterPrefix_(ctx *Prefix_Context) {}

// ExitPrefix_ is called when production prefix_ is exited.
func (s *Baseasm6502Listener) ExitPrefix_(ctx *Prefix_Context) {}

// EnterString_ is called when production string_ is entered.
func (s *Baseasm6502Listener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *Baseasm6502Listener) ExitString_(ctx *String_Context) {}

// EnterName is called when production name is entered.
func (s *Baseasm6502Listener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *Baseasm6502Listener) ExitName(ctx *NameContext) {}

// EnterNumber is called when production number is entered.
func (s *Baseasm6502Listener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *Baseasm6502Listener) ExitNumber(ctx *NumberContext) {}

// EnterComment is called when production comment is entered.
func (s *Baseasm6502Listener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *Baseasm6502Listener) ExitComment(ctx *CommentContext) {}

// EnterOpcode is called when production opcode is entered.
func (s *Baseasm6502Listener) EnterOpcode(ctx *OpcodeContext) {}

// ExitOpcode is called when production opcode is exited.
func (s *Baseasm6502Listener) ExitOpcode(ctx *OpcodeContext) {}
