// Code generated from CMake.g4 by ANTLR 4.7.2. DO NOT EDIT.

package cmake // CMake
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCMakeListener is a complete listener for a parse tree produced by CMakeParser.
type BaseCMakeListener struct{}

var _ CMakeListener = &BaseCMakeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCMakeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCMakeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCMakeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCMakeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseCMakeListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseCMakeListener) ExitFile(ctx *FileContext) {}

// EnterCommand_invocation is called when production command_invocation is entered.
func (s *BaseCMakeListener) EnterCommand_invocation(ctx *Command_invocationContext) {}

// ExitCommand_invocation is called when production command_invocation is exited.
func (s *BaseCMakeListener) ExitCommand_invocation(ctx *Command_invocationContext) {}

// EnterSingle_argument is called when production single_argument is entered.
func (s *BaseCMakeListener) EnterSingle_argument(ctx *Single_argumentContext) {}

// ExitSingle_argument is called when production single_argument is exited.
func (s *BaseCMakeListener) ExitSingle_argument(ctx *Single_argumentContext) {}

// EnterCompound_argument is called when production compound_argument is entered.
func (s *BaseCMakeListener) EnterCompound_argument(ctx *Compound_argumentContext) {}

// ExitCompound_argument is called when production compound_argument is exited.
func (s *BaseCMakeListener) ExitCompound_argument(ctx *Compound_argumentContext) {}
