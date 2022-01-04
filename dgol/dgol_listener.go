// Code generated from dgol.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dgol // dgol
import "github.com/antlr/antlr4/runtime/Go/antlr"

// dgolListener is a complete listener for a parse tree produced by dgolParser.
type dgolListener interface {
	antlr.ParseTreeListener

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterUsedeclaration is called when entering the usedeclaration production.
	EnterUsedeclaration(c *UsedeclarationContext)

	// EnterSubroutinedefinition is called when entering the subroutinedefinition production.
	EnterSubroutinedefinition(c *SubroutinedefinitionContext)

	// EnterProgramdefinition is called when entering the programdefinition production.
	EnterProgramdefinition(c *ProgramdefinitionContext)

	// EnterLibrarydefinition is called when entering the librarydefinition production.
	EnterLibrarydefinition(c *LibrarydefinitionContext)

	// EnterSubroutinedeclaration is called when entering the subroutinedeclaration production.
	EnterSubroutinedeclaration(c *SubroutinedeclarationContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterIdentifierorzero is called when entering the identifierorzero production.
	EnterIdentifierorzero(c *IdentifierorzeroContext)

	// EnterLetstatement is called when entering the letstatement production.
	EnterLetstatement(c *LetstatementContext)

	// EnterIfstatement is called when entering the ifstatement production.
	EnterIfstatement(c *IfstatementContext)

	// EnterIfhead is called when entering the ifhead production.
	EnterIfhead(c *IfheadContext)

	// EnterDostatement is called when entering the dostatement production.
	EnterDostatement(c *DostatementContext)

	// EnterCallstatement is called when entering the callstatement production.
	EnterCallstatement(c *CallstatementContext)

	// EnterReturnstatement is called when entering the returnstatement production.
	EnterReturnstatement(c *ReturnstatementContext)

	// EnterExitstatement is called when entering the exitstatement production.
	EnterExitstatement(c *ExitstatementContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitUsedeclaration is called when exiting the usedeclaration production.
	ExitUsedeclaration(c *UsedeclarationContext)

	// ExitSubroutinedefinition is called when exiting the subroutinedefinition production.
	ExitSubroutinedefinition(c *SubroutinedefinitionContext)

	// ExitProgramdefinition is called when exiting the programdefinition production.
	ExitProgramdefinition(c *ProgramdefinitionContext)

	// ExitLibrarydefinition is called when exiting the librarydefinition production.
	ExitLibrarydefinition(c *LibrarydefinitionContext)

	// ExitSubroutinedeclaration is called when exiting the subroutinedeclaration production.
	ExitSubroutinedeclaration(c *SubroutinedeclarationContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitIdentifierorzero is called when exiting the identifierorzero production.
	ExitIdentifierorzero(c *IdentifierorzeroContext)

	// ExitLetstatement is called when exiting the letstatement production.
	ExitLetstatement(c *LetstatementContext)

	// ExitIfstatement is called when exiting the ifstatement production.
	ExitIfstatement(c *IfstatementContext)

	// ExitIfhead is called when exiting the ifhead production.
	ExitIfhead(c *IfheadContext)

	// ExitDostatement is called when exiting the dostatement production.
	ExitDostatement(c *DostatementContext)

	// ExitCallstatement is called when exiting the callstatement production.
	ExitCallstatement(c *CallstatementContext)

	// ExitReturnstatement is called when exiting the returnstatement production.
	ExitReturnstatement(c *ReturnstatementContext)

	// ExitExitstatement is called when exiting the exitstatement production.
	ExitExitstatement(c *ExitstatementContext)
}
