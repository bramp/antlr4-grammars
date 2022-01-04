// Code generated from scotty.g4 by ANTLR 4.9.3. DO NOT EDIT.

package scotty // scotty
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasescottyListener is a complete listener for a parse tree produced by scottyParser.
type BasescottyListener struct{}

var _ scottyListener = &BasescottyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasescottyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasescottyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasescottyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasescottyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasescottyListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasescottyListener) ExitProg(ctx *ProgContext) {}

// EnterProgram_lines is called when production program_lines is entered.
func (s *BasescottyListener) EnterProgram_lines(ctx *Program_linesContext) {}

// ExitProgram_lines is called when production program_lines is exited.
func (s *BasescottyListener) ExitProgram_lines(ctx *Program_linesContext) {}

// EnterVar_assign is called when production var_assign is entered.
func (s *BasescottyListener) EnterVar_assign(ctx *Var_assignContext) {}

// ExitVar_assign is called when production var_assign is exited.
func (s *BasescottyListener) ExitVar_assign(ctx *Var_assignContext) {}

// EnterFn_def is called when production fn_def is entered.
func (s *BasescottyListener) EnterFn_def(ctx *Fn_defContext) {}

// ExitFn_def is called when production fn_def is exited.
func (s *BasescottyListener) ExitFn_def(ctx *Fn_defContext) {}

// EnterPrefix_exp is called when production prefix_exp is entered.
func (s *BasescottyListener) EnterPrefix_exp(ctx *Prefix_expContext) {}

// ExitPrefix_exp is called when production prefix_exp is exited.
func (s *BasescottyListener) ExitPrefix_exp(ctx *Prefix_expContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasescottyListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasescottyListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterId_tail is called when production id_tail is entered.
func (s *BasescottyListener) EnterId_tail(ctx *Id_tailContext) {}

// ExitId_tail is called when production id_tail is exited.
func (s *BasescottyListener) ExitId_tail(ctx *Id_tailContext) {}

// EnterNumber is called when production number is entered.
func (s *BasescottyListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasescottyListener) ExitNumber(ctx *NumberContext) {}

// EnterDigits is called when production digits is entered.
func (s *BasescottyListener) EnterDigits(ctx *DigitsContext) {}

// ExitDigits is called when production digits is exited.
func (s *BasescottyListener) ExitDigits(ctx *DigitsContext) {}
