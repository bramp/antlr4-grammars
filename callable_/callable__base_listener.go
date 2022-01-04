// Code generated from callable_.g4 by ANTLR 4.9.3. DO NOT EDIT.

package callable_ // callable_
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Basecallable_Listener is a complete listener for a parse tree produced by callable_Parser.
type Basecallable_Listener struct{}

var _ callable_Listener = &Basecallable_Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Basecallable_Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Basecallable_Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Basecallable_Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Basecallable_Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *Basecallable_Listener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *Basecallable_Listener) ExitProgram(ctx *ProgramContext) {}

// EnterLine is called when production line is entered.
func (s *Basecallable_Listener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *Basecallable_Listener) ExitLine(ctx *LineContext) {}

// EnterF_inner is called when production f_inner is entered.
func (s *Basecallable_Listener) EnterF_inner(ctx *F_innerContext) {}

// ExitF_inner is called when production f_inner is exited.
func (s *Basecallable_Listener) ExitF_inner(ctx *F_innerContext) {}

// EnterF_arg is called when production f_arg is entered.
func (s *Basecallable_Listener) EnterF_arg(ctx *F_argContext) {}

// ExitF_arg is called when production f_arg is exited.
func (s *Basecallable_Listener) ExitF_arg(ctx *F_argContext) {}
