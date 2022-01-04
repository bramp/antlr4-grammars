// Code generated from dgol.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dgol // dgol
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasedgolListener is a complete listener for a parse tree produced by dgolParser.
type BasedgolListener struct{}

var _ dgolListener = &BasedgolListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasedgolListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasedgolListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasedgolListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasedgolListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModule is called when production module is entered.
func (s *BasedgolListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BasedgolListener) ExitModule(ctx *ModuleContext) {}

// EnterUsedeclaration is called when production usedeclaration is entered.
func (s *BasedgolListener) EnterUsedeclaration(ctx *UsedeclarationContext) {}

// ExitUsedeclaration is called when production usedeclaration is exited.
func (s *BasedgolListener) ExitUsedeclaration(ctx *UsedeclarationContext) {}

// EnterSubroutinedefinition is called when production subroutinedefinition is entered.
func (s *BasedgolListener) EnterSubroutinedefinition(ctx *SubroutinedefinitionContext) {}

// ExitSubroutinedefinition is called when production subroutinedefinition is exited.
func (s *BasedgolListener) ExitSubroutinedefinition(ctx *SubroutinedefinitionContext) {}

// EnterProgramdefinition is called when production programdefinition is entered.
func (s *BasedgolListener) EnterProgramdefinition(ctx *ProgramdefinitionContext) {}

// ExitProgramdefinition is called when production programdefinition is exited.
func (s *BasedgolListener) ExitProgramdefinition(ctx *ProgramdefinitionContext) {}

// EnterLibrarydefinition is called when production librarydefinition is entered.
func (s *BasedgolListener) EnterLibrarydefinition(ctx *LibrarydefinitionContext) {}

// ExitLibrarydefinition is called when production librarydefinition is exited.
func (s *BasedgolListener) ExitLibrarydefinition(ctx *LibrarydefinitionContext) {}

// EnterSubroutinedeclaration is called when production subroutinedeclaration is entered.
func (s *BasedgolListener) EnterSubroutinedeclaration(ctx *SubroutinedeclarationContext) {}

// ExitSubroutinedeclaration is called when production subroutinedeclaration is exited.
func (s *BasedgolListener) ExitSubroutinedeclaration(ctx *SubroutinedeclarationContext) {}

// EnterStatements is called when production statements is entered.
func (s *BasedgolListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BasedgolListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasedgolListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasedgolListener) ExitStatement(ctx *StatementContext) {}

// EnterIdentifierorzero is called when production identifierorzero is entered.
func (s *BasedgolListener) EnterIdentifierorzero(ctx *IdentifierorzeroContext) {}

// ExitIdentifierorzero is called when production identifierorzero is exited.
func (s *BasedgolListener) ExitIdentifierorzero(ctx *IdentifierorzeroContext) {}

// EnterLetstatement is called when production letstatement is entered.
func (s *BasedgolListener) EnterLetstatement(ctx *LetstatementContext) {}

// ExitLetstatement is called when production letstatement is exited.
func (s *BasedgolListener) ExitLetstatement(ctx *LetstatementContext) {}

// EnterIfstatement is called when production ifstatement is entered.
func (s *BasedgolListener) EnterIfstatement(ctx *IfstatementContext) {}

// ExitIfstatement is called when production ifstatement is exited.
func (s *BasedgolListener) ExitIfstatement(ctx *IfstatementContext) {}

// EnterIfhead is called when production ifhead is entered.
func (s *BasedgolListener) EnterIfhead(ctx *IfheadContext) {}

// ExitIfhead is called when production ifhead is exited.
func (s *BasedgolListener) ExitIfhead(ctx *IfheadContext) {}

// EnterDostatement is called when production dostatement is entered.
func (s *BasedgolListener) EnterDostatement(ctx *DostatementContext) {}

// ExitDostatement is called when production dostatement is exited.
func (s *BasedgolListener) ExitDostatement(ctx *DostatementContext) {}

// EnterCallstatement is called when production callstatement is entered.
func (s *BasedgolListener) EnterCallstatement(ctx *CallstatementContext) {}

// ExitCallstatement is called when production callstatement is exited.
func (s *BasedgolListener) ExitCallstatement(ctx *CallstatementContext) {}

// EnterReturnstatement is called when production returnstatement is entered.
func (s *BasedgolListener) EnterReturnstatement(ctx *ReturnstatementContext) {}

// ExitReturnstatement is called when production returnstatement is exited.
func (s *BasedgolListener) ExitReturnstatement(ctx *ReturnstatementContext) {}

// EnterExitstatement is called when production exitstatement is entered.
func (s *BasedgolListener) EnterExitstatement(ctx *ExitstatementContext) {}

// ExitExitstatement is called when production exitstatement is exited.
func (s *BasedgolListener) ExitExitstatement(ctx *ExitstatementContext) {}
