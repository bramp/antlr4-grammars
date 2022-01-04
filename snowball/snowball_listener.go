// Code generated from snowball.g4 by ANTLR 4.9.3. DO NOT EDIT.

package snowball // snowball
import "github.com/antlr/antlr4/runtime/Go/antlr"

// snowballListener is a complete listener for a parse tree produced by snowballParser.
type snowballListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterP is called when entering the p production.
	EnterP(c *PContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterR_definition is called when entering the r_definition production.
	EnterR_definition(c *R_definitionContext)

	// EnterG_definition is called when entering the g_definition production.
	EnterG_definition(c *G_definitionContext)

	// EnterC is called when entering the c production.
	EnterC(c *CContext)

	// EnterI_command is called when entering the i_command production.
	EnterI_command(c *I_commandContext)

	// EnterS_command is called when entering the s_command production.
	EnterS_command(c *S_commandContext)

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterG is called when entering the g production.
	EnterG(c *GContext)

	// EnterS_name is called when entering the s_name production.
	EnterS_name(c *S_nameContext)

	// EnterI_name is called when entering the i_name production.
	EnterI_name(c *I_nameContext)

	// EnterB_name is called when entering the b_name production.
	EnterB_name(c *B_nameContext)

	// EnterR_name is called when entering the r_name production.
	EnterR_name(c *R_nameContext)

	// EnterG_name is called when entering the g_name production.
	EnterG_name(c *G_nameContext)

	// EnterAe is called when entering the ae production.
	EnterAe(c *AeContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitP is called when exiting the p production.
	ExitP(c *PContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitR_definition is called when exiting the r_definition production.
	ExitR_definition(c *R_definitionContext)

	// ExitG_definition is called when exiting the g_definition production.
	ExitG_definition(c *G_definitionContext)

	// ExitC is called when exiting the c production.
	ExitC(c *CContext)

	// ExitI_command is called when exiting the i_command production.
	ExitI_command(c *I_commandContext)

	// ExitS_command is called when exiting the s_command production.
	ExitS_command(c *S_commandContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitG is called when exiting the g production.
	ExitG(c *GContext)

	// ExitS_name is called when exiting the s_name production.
	ExitS_name(c *S_nameContext)

	// ExitI_name is called when exiting the i_name production.
	ExitI_name(c *I_nameContext)

	// ExitB_name is called when exiting the b_name production.
	ExitB_name(c *B_nameContext)

	// ExitR_name is called when exiting the r_name production.
	ExitR_name(c *R_nameContext)

	// ExitG_name is called when exiting the g_name production.
	ExitG_name(c *G_nameContext)

	// ExitAe is called when exiting the ae production.
	ExitAe(c *AeContext)
}
