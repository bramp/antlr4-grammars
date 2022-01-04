// Code generated from snowball.g4 by ANTLR 4.9.3. DO NOT EDIT.

package snowball // snowball
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesnowballListener is a complete listener for a parse tree produced by snowballParser.
type BasesnowballListener struct{}

var _ snowballListener = &BasesnowballListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesnowballListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesnowballListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesnowballListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesnowballListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasesnowballListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasesnowballListener) ExitProgram(ctx *ProgramContext) {}

// EnterP is called when production p is entered.
func (s *BasesnowballListener) EnterP(ctx *PContext) {}

// ExitP is called when production p is exited.
func (s *BasesnowballListener) ExitP(ctx *PContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BasesnowballListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BasesnowballListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterR_definition is called when production r_definition is entered.
func (s *BasesnowballListener) EnterR_definition(ctx *R_definitionContext) {}

// ExitR_definition is called when production r_definition is exited.
func (s *BasesnowballListener) ExitR_definition(ctx *R_definitionContext) {}

// EnterG_definition is called when production g_definition is entered.
func (s *BasesnowballListener) EnterG_definition(ctx *G_definitionContext) {}

// ExitG_definition is called when production g_definition is exited.
func (s *BasesnowballListener) ExitG_definition(ctx *G_definitionContext) {}

// EnterC is called when production c is entered.
func (s *BasesnowballListener) EnterC(ctx *CContext) {}

// ExitC is called when production c is exited.
func (s *BasesnowballListener) ExitC(ctx *CContext) {}

// EnterI_command is called when production i_command is entered.
func (s *BasesnowballListener) EnterI_command(ctx *I_commandContext) {}

// ExitI_command is called when production i_command is exited.
func (s *BasesnowballListener) ExitI_command(ctx *I_commandContext) {}

// EnterS_command is called when production s_command is entered.
func (s *BasesnowballListener) EnterS_command(ctx *S_commandContext) {}

// ExitS_command is called when production s_command is exited.
func (s *BasesnowballListener) ExitS_command(ctx *S_commandContext) {}

// EnterS is called when production s is entered.
func (s *BasesnowballListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BasesnowballListener) ExitS(ctx *SContext) {}

// EnterG is called when production g is entered.
func (s *BasesnowballListener) EnterG(ctx *GContext) {}

// ExitG is called when production g is exited.
func (s *BasesnowballListener) ExitG(ctx *GContext) {}

// EnterS_name is called when production s_name is entered.
func (s *BasesnowballListener) EnterS_name(ctx *S_nameContext) {}

// ExitS_name is called when production s_name is exited.
func (s *BasesnowballListener) ExitS_name(ctx *S_nameContext) {}

// EnterI_name is called when production i_name is entered.
func (s *BasesnowballListener) EnterI_name(ctx *I_nameContext) {}

// ExitI_name is called when production i_name is exited.
func (s *BasesnowballListener) ExitI_name(ctx *I_nameContext) {}

// EnterB_name is called when production b_name is entered.
func (s *BasesnowballListener) EnterB_name(ctx *B_nameContext) {}

// ExitB_name is called when production b_name is exited.
func (s *BasesnowballListener) ExitB_name(ctx *B_nameContext) {}

// EnterR_name is called when production r_name is entered.
func (s *BasesnowballListener) EnterR_name(ctx *R_nameContext) {}

// ExitR_name is called when production r_name is exited.
func (s *BasesnowballListener) ExitR_name(ctx *R_nameContext) {}

// EnterG_name is called when production g_name is entered.
func (s *BasesnowballListener) EnterG_name(ctx *G_nameContext) {}

// ExitG_name is called when production g_name is exited.
func (s *BasesnowballListener) ExitG_name(ctx *G_nameContext) {}

// EnterAe is called when production ae is entered.
func (s *BasesnowballListener) EnterAe(ctx *AeContext) {}

// ExitAe is called when production ae is exited.
func (s *BasesnowballListener) ExitAe(ctx *AeContext) {}
