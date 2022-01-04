// Code generated from ISL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package isl // ISL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseISLListener is a complete listener for a parse tree produced by ISLParser.
type BaseISLListener struct{}

var _ ISLListener = &BaseISLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseISLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseISLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseISLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseISLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseISLListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseISLListener) ExitProgram(ctx *ProgramContext) {}

// EnterDefOrExpr is called when production defOrExpr is entered.
func (s *BaseISLListener) EnterDefOrExpr(ctx *DefOrExprContext) {}

// ExitDefOrExpr is called when production defOrExpr is exited.
func (s *BaseISLListener) ExitDefOrExpr(ctx *DefOrExprContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseISLListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseISLListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseISLListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseISLListener) ExitExpr(ctx *ExprContext) {}

// EnterQuoted is called when production quoted is entered.
func (s *BaseISLListener) EnterQuoted(ctx *QuotedContext) {}

// ExitQuoted is called when production quoted is exited.
func (s *BaseISLListener) ExitQuoted(ctx *QuotedContext) {}

// EnterQuasiQuoted is called when production quasiQuoted is entered.
func (s *BaseISLListener) EnterQuasiQuoted(ctx *QuasiQuotedContext) {}

// ExitQuasiQuoted is called when production quasiQuoted is exited.
func (s *BaseISLListener) ExitQuasiQuoted(ctx *QuasiQuotedContext) {}

// EnterTestCase is called when production testCase is entered.
func (s *BaseISLListener) EnterTestCase(ctx *TestCaseContext) {}

// ExitTestCase is called when production testCase is exited.
func (s *BaseISLListener) ExitTestCase(ctx *TestCaseContext) {}

// EnterLibraryRequire is called when production libraryRequire is entered.
func (s *BaseISLListener) EnterLibraryRequire(ctx *LibraryRequireContext) {}

// ExitLibraryRequire is called when production libraryRequire is exited.
func (s *BaseISLListener) ExitLibraryRequire(ctx *LibraryRequireContext) {}

// EnterPkg is called when production pkg is entered.
func (s *BaseISLListener) EnterPkg(ctx *PkgContext) {}

// ExitPkg is called when production pkg is exited.
func (s *BaseISLListener) ExitPkg(ctx *PkgContext) {}

// EnterName is called when production name is entered.
func (s *BaseISLListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseISLListener) ExitName(ctx *NameContext) {}
